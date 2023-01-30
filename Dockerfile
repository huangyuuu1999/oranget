FROM 43.139.176.247/fruit_buckets/golang:latest

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /oranget
COPY src/app ./app

EXPOSE 8080
# CMD ["/oranget/app"]
ENTRYPOINT [ "./app" ]
