FROM golang:1.15

RUN go get github.com/ucef29/statExporter
WORKDIR /go/src/github.com/ucef29/statExporter
RUN go build -o /go/statExporter .

