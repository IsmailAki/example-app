### Stage 1
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/example-app

RUN apk add upx
RUN upx --ultra-brute /app/example-app

### Stage 2
FROM scratch
COPY --from=builder /app/example-app /bin/example-app
ENTRYPOINT ["/bin/example-app"]
