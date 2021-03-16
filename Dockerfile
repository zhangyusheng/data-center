FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zhangyusheng/data-center
COPY . $GOPATH/src/github.com/zhangyusheng/data-center
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./data-center"]
