FROM golang:1.13-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /build
COPY . .
RUN go build -a -o "/usr/local/bin/relay-rainbow-step-bubble-message-send" "./sendmsgtobubble.go"

FROM alpine:3
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder "/usr/local/bin/relay-rainbow-step-bubble-message-send" "/usr/local/bin/relay-rainbow-step-bubble-message-send"
ENTRYPOINT []
CMD ["/usr/local/bin/relay-rainbow-step-bubble-message-send"]

LABEL "org.opencontainers.image.title"="Send a message to a Rainbow bubble"
LABEL "org.opencontainers.image.description"="This step sends a message to a Rainbow bubble."
LABEL "sh.relay.sdk.version"="v1"
