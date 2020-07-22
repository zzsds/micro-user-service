FROM golang:1.13-alpine as builder
ARG path
WORKDIR /var/app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -o app -a -installsuffix cgo main.go plugin.go

FROM alpine
# 设置时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo '$TZ' > /etc/timezone

RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories && \
    apk --no-cache add ca-certificates
ENV CFG_ENDPOINT=acm.aliyun.com:8080
ENV CFG_GROUP=prod
WORKDIR /app
COPY --from=builder /var/app/app .
ENTRYPOINT [ "./app" ]
