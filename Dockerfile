FROM golang:1.21.4-alpine AS build
WORKDIR /go/src/app

COPY go.* ./
RUN go mod tidy

COPY . .

RUN go build -o bin/main cmd/main.go

FROM alpine:3.15

COPY --from=build /go/src/app /go/src/app

WORKDIR /go/src/app

ADD https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz /usr/local/bin/
RUN tar -C /usr/local/bin -xzvf /usr/local/bin/dockerize-linux-amd64-v0.6.1.tar.gz
RUN rm /usr/local/bin/dockerize-linux-amd64-v0.6.1.tar.gz

EXPOSE 8080

CMD ["dockerize", "-wait", "tcp://mysql:3306", "-timeout", "30s", "./bin/main"]


