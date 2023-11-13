FROM golang:1.20.2-alpine3.17 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod download
RUN go build -o server cmd/api/*.go

FROM alpine:latest

WORKDIR /api

RUN apk add tzdata

COPY --from=builder /app/server /server

EXPOSE 8000

CMD [ "/server", "--port", "8000"]
