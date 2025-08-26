FROM golang:1.25.0-trixie AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o worker ./worker

FROM alpine:3.22.1

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/worker /worker

ENTRYPOINT ["/worker"]

