FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod init example.com/app

RUN go build -o main .

CMD ["./main"]
