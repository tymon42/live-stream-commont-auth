# live-stream-comment-auth  

English | [简体中文](README_zh.md)

[中文介绍文章](https://www.bilibili.com/read/cv19545136)

Originally, I wanted to write a third-party live interactive tool for Bilibili, but was troubled by identity verification and login design, so I came up with this idea. Me and [@Cunoe](https://github.com/CUNOE) have already implemented a basic version, which is currently working properly. Now I am developing a more universal open source version and will provide paid services (of course, there will also be a free trial period).
## What's this?  
This is a designed authorization solution for third-party applications on live streaming platforms (such as Bilibili or YouTube). You can easily use RESTful API services to authorize and generate JWT tokens for your own applications (or barrage games). Only users registered on specific live streaming platforms can pass the authorization process.

We believe that this type of message is already secure enough for login operations.

Please check how it works to get more detailed information.

Check [*How it works*](docs/how-it-works.md) for more details.

<!-- ## Package Usage  
Use this package as go mod:  
Install:  
```
go get github.com/tymon42/live-stream-comment-auth
```
In program:  
```
import "github.com/tymon42/live-stream-comment-auth/vcode"

new_vcode := vcode.GenBiliVCodeWithExtraInfo("<UUID>", "<UID>", "<TIME_NOW>")
``` -->

## Use the official Saas service
More details, please check [*Saas documantation*](docs/saas.md)

<!-- ## Authing as a Service (Bilibili)
### Start via Docker
```
docker build -t bili-danmu-auth-api -f Dockerfile .

docker run \
  -p 8888:8888 \
  -v ${PWD}/data:/data \
  -e REPLICA_URL=<YOUR_DATA_BASE_BACKUP_URL> \
  bili-danmu-auth-api
```
#### Start worker
```
go run bili-danmu-auth/worker/auther/main.go -api "http://127.0.0.1:8888"  -r <BILI_ROOM_ID> -p "" -l 6 -k "IAMAWORKER" &
``` -->

<!-- #### Start Swagger web UI
```
docker run --platform linux/amd64 --rm -p 8083:8080 -e SWAGGER_JSON_URL=/swagger/bili-danmu-auth.json -e SWAGGER_JSON=/foo/bili-danmu-auth.json -v $PWD/bili-danmu-auth/api:/usr/share/nginx/html/swagger swaggerapi/swagger-ui
``` -->

### Youtube
WIP
