FROM golang:alpine

COPY . /go/src/app
WORKDIR /go/src/app

RUN go mod init helloworld
RUN go mod tidy
RUN go build -o helloworld

CMD ["./helloworld"]
EXPOSE 6111