FROM 43.139.176.247/fruit_buckets/golang:latest

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /oranget
ADD ./go.mod .
ADD ./go.sum .
COPY src ./src

RUN ["go", "build", "./src/main.go"]

EXPOSE 8080
# CMD ["/oranget/app"]
ENTRYPOINT [ "./main" ]
