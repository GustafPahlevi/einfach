
FROM golang:latest

LABEL maintainer="Gustaf Pahlevi"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod vendor

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main", "serve"]