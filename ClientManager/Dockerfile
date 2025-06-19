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

EXPOSE 8081

CMD ["./main"]