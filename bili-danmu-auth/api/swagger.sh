#!/bin/bash
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
goctl api plugin -plugin goctl-swagger="swagger -filename bili-danmu-auth.json" -api bili-danmu-auth.api -dir .