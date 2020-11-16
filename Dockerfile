FROM golang:latest

ENV GOPROXY https://proxy.golang.org,direct
WORKDIR $GOPATH/src/github.com/ROGGER1808/go-gin-example
COPY . $GOPATH/src/github.com/ROGGER1808/go-gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]

