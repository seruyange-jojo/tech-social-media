FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]