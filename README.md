# [WIP] live-stream-commont-auth
I came up with this idea when I wanted to build a Bilibili third party live-interactive tool. And [@Cunoe](https://github.com/CUNOE) and I have already implemented a draft version and it is working well on our tool.  
Now I am working on open source version which is gonna be more universal.   
# What's this?  
It's a kind of authorization solution design for 3rd party app in live streaming platform like Bilibili or Youtube and so on.  

You can easily setup a self-host auth service and auth and generate JWT token for your own app. Only registered user on the specific live streaming platform can pass thourgh the auth process.  

# How it works?
The main idea of the live-stream-comment-auth is capturing infomation that post by registered users on platform(like bilibili or youtube) only. The infomation must has the uid that represents the particular user.  

For example, danmu in bilibili live has an excellent data structure for the identity authorzing. It includes bili user's uid and a 20 chars space to convey a message. If a user wants to login to a third party platform, we could sent the user a verification code in 20 chars and the user should send it to a chose bilibili live room as a danmu. So we shall get a danmu message via a bili's danmu WS connection. And the message contains both the buid and the verification code.  

We believe that kind of message is safe enough for a login action.  

## Package Usage  
Use this package as go mod:  
Install:  
```
go get github.com/tymon42/live-stream-comment-auth
```
In program:  
```
import "github.com/tymon42/live-stream-comment-auth/vcode"

new_vcode := vcode.GenBiliVCodeWithExtraInfo("<UUID>", "<UID>", "<TIME_NOW>")
```

## Example  
> Authing As a Service
### Bilibili
#### Start server
```
docker build -f bili-danmu-auth/api/Dockerfile -t bili-danmu-auth-server:latest .
docker run --rm -it -p 8888:8888 bili-danmu-auth-server:latest
```
#### Start worker
```
go run bili-danmu-auth/worker/main.go -r <room_id>
```

#### Start Swagger web UI
```
docker run --platform linux/amd64 --rm -p 8083:8080 -e SWAGGER_JSON_URL=/swagger/bili-danmu-auth.json -e SWAGGER_JSON=/foo/bili-danmu-auth.json -v $PWD/bili-danmu-auth/api:/usr/share/nginx/html/swagger swaggerapi/swagger-ui
```

### Youtube
WIP
