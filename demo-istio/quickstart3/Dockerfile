# Build stage
FROM --platform=$TARGETPLATFORM golang:1.19.4-alpine3.17 AS builder
ARG TARGETARCH
ARG TARGETOS
WORKDIR /app
COPY pb/ pb/
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
RUN go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o main main.go

# Run stage
FROM --platform=$TARGETPLATFORM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8090
CMD ["/app/main"]
ENTRYPOINT [ "/app/main" ]