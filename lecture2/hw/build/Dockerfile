FROM golang:1.20.2-alpine3.16 AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./
RUN go mod download
RUN GOOS=linux go build -o ./app ./cmd/

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /app/ .
#COPY --from=0 /app/docs ./docs

EXPOSE 8000

CMD ["./app"]