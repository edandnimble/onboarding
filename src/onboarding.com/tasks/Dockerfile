FROM golang:1.17 AS builder
# needed for alpine to run code
env CGO_ENABLED=0
COPY . /build
WORKDIR /build/tasks
RUN go build -o tasks.bin ./main/main.go

FROM alpine
COPY --from=builder /build/tasks/tasks.bin /tasks
cmd ["/tasks"]
