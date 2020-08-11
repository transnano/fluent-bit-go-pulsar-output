# fluent-bit-go-pulsar-output-sample

Fluent-Bit go apache pulsar output plugin

## Build a Go Plugin

To build the code above, use the following line:

```sh
$ go build -buildmode=c-shared -o out_pulsar.so .
```

Once built, a shared library called `out_pulsar.so` will be available. It's really important to double check the final .so file is what we expect. Doing a ldd over the library we should see something similar to this:

```sh
$ ldd out_pulsar.so
        linux-vdso.so.1 (0x00007ffe445fa000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fdd13af6000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fdd13935000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fdd14848000)
```

## Run Fluent Bit with the new plugin

```sh
$ bin/fluent-bit -e /path/to/out_pulsar.so -i cpu -o flb-go-pulsar
```


```sh

$ docker-compose up -d .
$ 
$ 

$ go build -buildmode=c-shared -o out_pulsar.so .
$ sudo docker run -it --rm -p 127.0.0.1:24224:24224 fluent/fluent-bit:1.5.2 /fluent-bit/bin/fluent-bit -i forward -o stdout -p format=json_lines -f 1

$ sudo docker run --rm fluent/fluent-bit:1.5.2 /fluent-bit/bin/fluent-bit -i cpu -o stdout -p format=json_lines -f 1

$ sudo docker run -it --rm -v $(pwd):/pulsar -p 127.0.0.1:24224:24224 fluent/fluent-bit:1.5.2 /fluent-bit/bin/fluent-bit -i cpu -c /pulsar/example/fluent.conf -e /pulsar/out_pulsar.so -o stdout -p format=json_lines -f 1


$ sudo docker build -t my-pulsar .

$ sudo docker run -it --rm -v $(pwd)/example:/usr/local/src my-pulsar:latest /fluent-bit/bin/fluent-bit -i cpu -c /usr/local/fluent.conf -e /fluent-bit/bin/out_pulsar.so -i cpu -o flb-go-pulsar -p format=json_lines -f 1
$ sudo docker run --rm -v $(pwd)/example:/usr/local/src my-pulsar:latest /fluent-bit/bin/fluent-bit -i cpu -e /fluent-bit/bin/out_pulsar.so -i cpu -o flb-go-pulsar -p format=json_lines -f 1


# 2020/08/09
$ sudo docker build -t my-pulsar .
$ sudo docker run --rm \
-e FLB_GO_PULSAR_BROKER_SERVICE_URL=pulsar://localhost:6650 \
-e FLB_GO_PULSAR_TENNANT=pulsar \
-e FLB_GO_PULSAR_NAMESPACE=default \
-e FLB_GO_PULSAR_TOPIC=test \
-e FLB_GO_PULSAR_TLS_ENABLED=true \
-e FLB_GO_PULSAR_TLS_TRUST_CERTS_FILE_PATH=file://path/to/cert \
-e FLB_GO_PULSAR_TLS_ALLOW_INSECURE_CONNECTION=true \
my-pulsar:latest /fluent-bit/bin/fluent-bit -i cpu -c /fluent-bit/etc/fluent-bit.conf -e /fluent-bit/bin/out_pulsar.so -i cpu -o flb-go-pulsar -p format=json_lines -f 1

$ sudo docker run --rm --env-file .env my-pulsar:latest /fluent-bit/bin/fluent-bit -i cpu -c /fluent-bit/etc/fluent-bit.conf -e /fluent-bit/bin/out_pulsar.so -i cpu -o flb-go-pulsar -p format=json_lines -f 1

$ sudo docker run --rm my-pulsar:latest /fluent-bit/bin/fluent-bit -i cpu -c /fluent-bit/etc/fluent-bit.conf -e /fluent-bit/bin/out_pulsar.so -i cpu -o flb-go-pulsar -p format=json_lines -f 1
```
