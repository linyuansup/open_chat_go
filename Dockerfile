FROM golang:latest as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /
COPY . .
RUN GOOS=linux GOARCH=amd64
RUN go test -timeout 30s -run ^TestRunner$ opChat -v
RUN go build
FROM alpine as runner
WORKDIR /
COPY --from=builder /opChat .
COPY --from=builder /storage .
EXPOSE 80
ENTRYPOINT ["ls && ./opChat"]