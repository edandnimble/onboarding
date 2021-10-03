FROM golang:1.17 AS builder
# needed for alpine to run code
env CGO_ENABLED=0
COPY . /build
WORKDIR /build/guesser
RUN go build -o guesser.bin ./main/main.go

FROM alpine
COPY --from=builder /build/guesser/guesser.bin /guesser
cmd ["/guesser"]
