
FROM golang:1.25.4-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o godocker ./cmd/api/main.go 

EXPOSE 8080

CMD ["./godocker"]