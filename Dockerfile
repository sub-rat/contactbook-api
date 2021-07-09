FROM golang:latest

WORKDIR /go/src/github.com/sub-rat/contactbook-api
COPY . .
RUN go build -o /usr/bin/contactbook-api cmd/api/main.go

CMD ["contactbook-api"]
EXPOSE 8080