# Build stage
FROM golang:alpine AS build-env

COPY . /go/src/github.com/Ullaakut/goneypot
WORKDIR /go/src/github.com/Ullaakut/goneypot/cmd

RUN go version
RUN go build -o goneypot

# Final stage
FROM alpine

WORKDIR /app
COPY --from=build-env /go/src/github.com/Ullaakut/goneypot/cmd/ /app/
COPY --from=build-env /go/src/github.com/Ullaakut/goneypot/config.yaml /etc/goneypot/config.yaml

ENTRYPOINT ["/app/goneypot"]
