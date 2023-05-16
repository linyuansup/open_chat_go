FROM golang as builder
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
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
EXPOSE 80
ENTRYPOINT ["./opChat"]