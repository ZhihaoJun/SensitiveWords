FROM golang:1.7

COPY src/sensitive/ /go/src/sensitive/
COPY sensitives.txt /sensitives.txt
RUN go build -o /main src/sensitive/main.go
ENTRYPOINT ["/main"]
