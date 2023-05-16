FROM golang:latest as builder
RUN apk --no-cache add tzdata
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0
WORKDIR /
COPY . .
RUN go test -timeout 30s -run ^TestRunner$ opChat -v
RUN go build .
FROM scratch
WORKDIR /
COPY --from=builder ./opChat .
COPY --from=builder ./storage ./storage
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
EXPOSE 80
ENTRYPOINT ["./opChat"]