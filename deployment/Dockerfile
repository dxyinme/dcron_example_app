FROM golang:alpine3.17 AS builder
LABEL maintainer="dxyinme@outlook.com"
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN cd app && go mod tidy
RUN go build -o /app/app app/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/app /app/app
EXPOSE 8080
CMD [ "./app" ]