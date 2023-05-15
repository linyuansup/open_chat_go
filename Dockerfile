FROM golang:latest as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /
COPY . .
RUN GOOS=linux GOARCH=amd64 go build
RUN mkdir publish && cp opChat publish && \
    cp -r storage publish
FROM alpine
WORKDIR /
COPY --from=builder /app/publish .
EXPOSE 80
ENTRYPOINT ["./opChat"]