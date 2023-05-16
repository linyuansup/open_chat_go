FROM golang:latest as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0
WORKDIR /
COPY . .
RUN go test -timeout 30s -run ^TestRunner$ opChat -v
RUN go build .
FROM alpine
WORKDIR /
COPY --from=builder ./opChat .
COPY --from=builder ./storage ./storage
EXPOSE 80
ENTRYPOINT ["./opChat"]