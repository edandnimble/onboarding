FROM golang:1.17 AS builder
# needed for alpine to run code
env CGO_ENABLED=0
COPY . /build
WORKDIR /build/api
RUN go build -o api.bin ./main/main.go

FROM alpine
COPY --from=builder /build/api/api.bin /api
cmd ["/api"]
