FROM golang:1.15-alpine3.12 AS build

WORKDIR /yt-relay

COPY . .

RUN apk add --update --no-cache make && \
    go get ./... && \
    make all

FROM alpine:3.12

WORKDIR /yt-relay

COPY --from=build /yt-relay/bin/ .
COPY --from=build /yt-relay/config.yml .

EXPOSE 8080
CMD ["./yt-relay", "serve", "-config", "./config.yml", "-address", "0.0.0.0", "-port", "8080"]
