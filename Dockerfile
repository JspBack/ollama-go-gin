FROM golang:1.23-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://proxy.golang.org,direct

RUN apk add --no-cache make curl

COPY . .
RUN go mod tidy && \
    go install github.com/swaggo/swag/cmd/swag@latest

RUN make swag && \
    make prod-lint && \
    make prod

FROM alpine:latest 

WORKDIR /app

COPY --from=builder /app/server /server

ARG PORT=8000

EXPOSE $PORT

CMD ["/server"]
