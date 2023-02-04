FROM golang:1.20.0-buster AS build

WORKDIR /app

COPY . /app
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o notifications-cli .

FROM alpine:3.17.1
WORKDIR /app
COPY --from=build /app/notifications-cli ./
COPY --from=build /app/config.yaml ./

ENTRYPOINT ["./notifications-cli"]