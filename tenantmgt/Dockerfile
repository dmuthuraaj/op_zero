# builder image
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main cmd/main.go

# generate clean, final image for end users
FROM alpine:latest
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/main .

CMD ["./main"]