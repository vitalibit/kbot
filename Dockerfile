# Build stage
FROM golang:1.20 as builder
WORKDIR /go/src/app
COPY . .
RUN make build

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/app/kbot .
ENTRYPOINT ["./kbot", "start"]