FROM golang:1.14.7 as builder

COPY go.mod /build/
COPY go.sum /build/
COPY out_pulsar.go /build/
RUN cd /build/ && go build -buildmode=c-shared -o out_pulsar.so .

FROM fluent/fluent-bit:1.5.2 as fluent-bit
USER root

COPY --from=builder /build/out_pulsar.so /fluent-bit/bin/
COPY --from=builder /build/out_pulsar.h /fluent-bit/bin/

CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf", "-e", "/fluent-bit/bin/out_pulsar.so"]
