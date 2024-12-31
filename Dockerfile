# Stage 1: Build stage
FROM golang:1.22.9-alpine3.19 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go install github.com/google/wire/cmd/wire@latest


RUN cd internal/ && wire
RUN GOOS=linux GOARCH=amd64 go build -tags musl --ldflags "-extldflags '-static' -s -w" -o /build/app-bin ./cmd/main.go

# Final stage: Run stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /build/app-bin /build/app-bin
COPY --from=builder /app .

EXPOSE 5000
ENTRYPOINT ["/build/app-bin"]
