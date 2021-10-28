FROM golang:1.17.2 as builder

WORKDIR /build/
COPY go.mod ./
COPY go.sum ./
COPY out_pulsar.go ./
COPY *.conf ./

RUN go build -buildmode=c-shared -o out_pulsar.so .

FROM fluent/fluent-bit:1.8.9 as fluent-bit
# hadolint ignore=DL3002
USER root

COPY --from=builder /build/out_pulsar.so /fluent-bit/bin/
COPY --from=builder /build/out_pulsar.h /fluent-bit/bin/
COPY --from=builder /build/*.conf /fluent-bit/etc/

CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf"]
