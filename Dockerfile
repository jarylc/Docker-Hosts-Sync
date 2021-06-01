FROM golang:alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s"

FROM scratch
ENV INTERVAL=60 \
    EXIT_RESET=1
COPY --from=builder /build/docker-hosts-sync /usr/bin/docker-hosts-sync
ENTRYPOINT ["/usr/bin/docker-hosts-sync"]
