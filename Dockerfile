FROM golang:1.14.4 AS builder
WORKDIR /app
COPY src .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o emoji-udp-server .

FROM scratch
COPY --from=builder /app/emoji-udp-server /
ENV EMOJI_PORT 54321
CMD ["/emoji-udp-server", "-n", "3", "-s", "','"]
