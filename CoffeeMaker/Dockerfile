FROM golang:alpine AS builder

WORKDIR /CoffeeMaker

COPY go.mod go.sum ./
RUN go mod download

COPY  . .

WORKDIR /CoffeeMaker/cmd

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /CoffeeMaker/cmd/main .

COPY --from=builder /CoffeeMaker/configs/config.yaml ./config.yaml

CMD ["./main"]