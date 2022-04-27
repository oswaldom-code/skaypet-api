#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -v -d ./...
RUN go build -o bin/skaypet-api main.go
CMD ["/go/src/app/bin/skaypet-api", "server"]
EXPOSE 9000