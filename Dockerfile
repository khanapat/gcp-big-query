FROM golang:1.15-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o goapp

FROM alpine:3.12

WORKDIR /app

COPY config/config.yaml config/config.yaml
COPY --from=build /app/goapp .

CMD ["./goapp"]

