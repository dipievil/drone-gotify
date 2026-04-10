# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod ./
# Copy go.sum if it exists
COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o drone-gotify .

# Final stage
FROM alpine:latest

RUN apk update && \
    apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

COPY --from=builder /app/drone-gotify /bin/

ENTRYPOINT ["/bin/drone-gotify"]
