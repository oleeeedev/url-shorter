FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o url-shorter

ENV PORT=8080

EXPOSE 8080

CMD ["./url-shorter"]
