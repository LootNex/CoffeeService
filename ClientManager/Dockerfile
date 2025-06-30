FROM golang:alpine AS builder

WORKDIR /ClientManager

COPY go.mod go.sum ./
RUN go mod download

COPY  . .

WORKDIR /ClientManager/cmd

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /ClientManager/cmd/main .

COPY --from=builder /ClientManager/configs/config.yaml ./config.yaml

EXPOSE 8081

CMD ["./main"]