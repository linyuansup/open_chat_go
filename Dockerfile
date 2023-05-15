FROM golang:latest as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /
COPY . .
RUN GOOS=linux GOARCH=amd64
RUN go test -timeout 30s -run ^TestRunner$ opChat -v
RUN go build
RUN ls
RUN mkdir publish && cp opChat publish && \
    cp -r storage publish
FROM alpine
WORKDIR /
COPY --from=builder /publish .
EXPOSE 80
ENTRYPOINT ["./opChat"]