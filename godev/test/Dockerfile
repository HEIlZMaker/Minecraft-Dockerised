FROM golang:1.20-alpine

WORKDIR /app

COPY ./test.go ./test.go

RUN go build test.go 

CMD ["./test"]
