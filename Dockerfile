
FROM golang:latest

RUN go get github.com/gin-gonic/gin && \
	go get github.com/thiduzz/todo

RUN mkdir /go/public

# Getting a simple example
RUN mv /go/src/github.com/thiduzz/todo/main.go /go/public/

CMD go run /go/public/main.go

EXPOSE 8080