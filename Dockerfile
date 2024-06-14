#build stage
FROM golang:alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
#Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY --from=builder /app/migrate ./migrate
COPY db/migration ./migration
COPY start.sh .
COPY wait-for-it.sh .
RUN apk add --no-cache bash
EXPOSE 8080
ENTRYPOINT [ "./start.sh" ]
