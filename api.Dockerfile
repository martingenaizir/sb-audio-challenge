FROM golang:1.22-alpine AS build
LABEL maintainer="Martin Genaizir <martingenaizir@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /root

RUN apk --no-cache add ffmpeg
COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
