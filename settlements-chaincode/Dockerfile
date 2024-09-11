FROM golang:1.21-alpine

WORKDIR /go/src/chaincode

COPY . .

RUN go build -mod=vendor -o chaincode

CMD ["./chaincode"]
