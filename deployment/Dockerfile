FROM golang:alpine3.17 AS builder
LABEL maintainer="dxyinme@outlook.com"
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN go build -o /app/app app/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /build/app/etc/app.yaml /app/etc/app.yaml
EXPOSE 8080

# You should set these environment variables when use this image.
# APP_DCRON_SERVICENAME: {str: dcron service name. default: "dcronapp"}
# APP_ENABLEREPORTER: {str: "true" or "false". default: "false"}
# APP_INNERCALL_CHANNEL: {str: app innercall channel. default: "dcronapp-channel"}
# APP_INNERCALL_REDIS_ADDR: {str: app innercall redis addr. default: "localhost:2379"}
# APP_INNERCALL_REDIS_DB: {int: app innercall redis db (should different with dcron redis). default: 2}
# APP_INNERCALL_REDIS_PASSWORD: {str: redis innercall password. default: ""}
# APP_MYSQL_ADDR: {str: mysql addr. default: ""}
# APP_MYSQL_PASSWORD: {str: mysql password. default: ""}
# APP_MYSQL_USER: {str: mysql user. defualt: ""}
# APP_PORT: {int: app serve port. default: 8080}
# APP_REDIS_ADDR: {str: app dcron redis addr. default: localhost:2379}
# APP_REDIS_DB: {int: app dcron redis db. default: 0}
# APP_REDIS_PASSWORD: {str: app redis password. default: ""}

CMD [ "./app", "-mode", "FromEnv" ]