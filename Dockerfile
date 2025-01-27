FROM golang:1.22-alpine3.18
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rss-aggregator ./cmd/rss-aggregator
EXPOSE 8000
CMD ["./rss-aggregator"]