FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY bili-danmu-auth/api/etc /usr/local/bin/etc
RUN --mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg \
    go build -ldflags='-s -w -extldflags "-static"' -tags osusergo,netgo,sqlite_omit_load_extensio -o /usr/local/bin/danmu-auth-api bili-danmu-auth/api/danmu-auth.go
# RUN go build -ldflags="-s -w" -o /app/danmu-auth-api bili-danmu-auth/api/danmu-auth.go

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz


# FROM scratch
FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /usr/local/bin/danmu-auth-api /usr/local/bin/danmu-auth-api
COPY --from=builder /usr/local/bin/etc /usr/local/bin/etc
COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream

RUN apk add bash

EXPOSE 8888

COPY etc/litestream.yml /etc/litestream.yml
COPY script/run.sh /script/run.sh

CMD [ "/script/run.sh" ]

# CMD ["./danmu-auth", "-f", "etc/danmu-auth.yaml"]
