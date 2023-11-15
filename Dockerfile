FROM golang:1.21.4-alpine3.17 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o server cmd/api/*.go

FROM alpine:latest

WORKDIR /api

COPY --from=builder /app/server /server
COPY files files/

EXPOSE 8000

CMD [ "/server"]
