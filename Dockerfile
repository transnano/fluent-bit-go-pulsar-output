FROM golang:1.15.0 as builder

COPY go.mod /build/
COPY go.sum /build/
COPY out_pulsar.go /build/
COPY plugins.conf /build/
COPY example/fluent.conf /build/
COPY sp-samples-1k.log /build/
RUN cd /build/ && go build -buildmode=c-shared -o out_pulsar.so .

FROM fluent/fluent-bit:1.5.3 as fluent-bit
USER root

COPY --from=builder /build/out_pulsar.so /fluent-bit/bin/
COPY --from=builder /build/out_pulsar.h /fluent-bit/bin/
COPY --from=builder /build/plugins.conf /fluent-bit/etc/
COPY --from=builder /build/fluent.conf /fluent-bit/etc/fluent-bit.conf
COPY --from=builder /build/sp-samples-1k.log /fluent-bit/etc/sp-samples-1k.log

CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf"]
