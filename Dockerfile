FROM golang:1.21.4-alpine AS build
WORKDIR /go/src/app

COPY . .

RUN go mod tidy

RUN go build -o cmd/bin/main cmd/main.go

RUN go install github.com/rubenv/sql-migrate/...@latest
ADD https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz /usr/local/bin/

FROM alpine:3.15

WORKDIR /go/src/app

COPY --from=build  /go/src/app/cmd/bin ./cmd/bin
COPY --from=build  /go/src/app/uploads ./uploads
COPY --from=build  /go/src/app/config.yml .
COPY --from=build  /go/src/app/datasource/mysql/migrations ./datasource/mysql/migrations
COPY --from=build  /go/src/app/datasource/mysql/dbconfig.yml ./datasource/mysql/dbconfig.yml

COPY --from=build /go/bin/sql-migrate /usr/local/bin/sql-migrate
COPY --from=build /usr/local/bin /usr/local/bin

RUN tar -C /usr/local/bin -xzvf /usr/local/bin/dockerize-linux-amd64-v0.6.1.tar.gz
RUN rm /usr/local/bin/dockerize-linux-amd64-v0.6.1.tar.gz


EXPOSE 8080

CMD ["dockerize", \
    "-wait", "tcp://mysql:3306", \
    "-timeout", "30s", \ 
    "sh", "-c", \
    "sql-migrate up -env=production -config=/go/src/app/datasource/mysql/dbconfig.yml && /go/src/app/cmd/bin/main"]


