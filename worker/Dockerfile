FROM golang:latest AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -mod=readonly -v -o worker ./worker

FROM alpine:3.22.1

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/worker/worker /worker

CMD ["/worker"]
