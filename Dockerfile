FROM golang:1.17.0-alpine3.14

RUN mkdir /app
WORKDIR /app
ADD . /app

EXPOSE 8080/tcp

RUN go build -o main ./server

CMD ["/app/main"]