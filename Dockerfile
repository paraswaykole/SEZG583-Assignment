FROM golang:1.22.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o server .
EXPOSE 3000
CMD ["/app/server"]
