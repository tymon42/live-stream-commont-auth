# live-stream-comment-auth  
I came up with this idea when I wanted to build a Bilibili third party live-interactive tool. And [@Cunoe](https://github.com/CUNOE) and I have already implemented a draft version and it is working well on our tool.  
Now I am working on open source version which is gonna be more universal. And I will host a services that require payment(but of couse there will be a free tier) as well.
## What's this?  
It's a kind of authorization solution design for 3rd party app in live streaming platform like Bilibili or Youtube and so on.  

You can easily use the RESTFUL API service to auth and generate JWT token for your own app (or danmu-game). Only registered user on the specific live streaming platform can pass thourgh the auth process.  


We believe that kind of message is safe enough for a login action.  

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

## Test
