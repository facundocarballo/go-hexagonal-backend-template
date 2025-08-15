# Build the binary
FROM golang:1.24-alpine AS builder

COPY go.mod /opt/app/
COPY go.sum /opt/app/

WORKDIR /opt/app

RUN go mod download

COPY src /opt/app/src

# Compile the binary
RUN go build -o server ./src

# Final minimalist image
FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /opt/app/server .

# Start the sever
CMD ["./server"]