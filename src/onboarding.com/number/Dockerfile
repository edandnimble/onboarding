FROM golang:1.17 AS builder
# needed for alpine to run code
env CGO_ENABLED=0
COPY . /build
WORKDIR /build/number
RUN go mod download
RUN go build -o number.bin ./main/main.go

FROM alpine
COPY --from=builder /build/number/number.bin /number
cmd ["/number"]
