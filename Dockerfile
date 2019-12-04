FROM golang:1.13.3 as builder

ENV GOPROXY="https://goproxy.cn,direct"

COPY . /api-test-tool

WORKDIR /api-test-tool

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/apitest -mod=vendor -ldflags '-s -w' ./main.go

FROM debian:8

COPY --from=builder /api-test-tool/bin/apitest /usr/local/bin/apitest

CMD ["app"]
