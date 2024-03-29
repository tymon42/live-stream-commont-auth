#!/bin/bash
set -e

# Restore the database if it does not already exist.
if [ -f /data/danmu-auth-api.db ]; then
	echo "Database already exists, skipping restore"
else
	echo "No database found, restoring from replica if exists"
	litestream restore -v -if-replica-exists -o /data/danmu-auth-api.db "${REPLICA_URL}"
fi

# Run litestream with your app as the subprocess.
exec litestream replicate -exec "/usr/local/bin/danmu-auth-api -f /usr/local/bin/etc/danmu-auth.yaml -db-path /data/danmu-auth-api.db"
