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
$ bin/fluent-bit -e /path/to/out_pulsar.so -i cpu -o pulsar-go -p plugin_conf1=value1 ...
```


```sh
$ sudo docker-compose up -d pulsar
# wait a minute...
$ sudo docker-compose up --build fluentbit
```

Run Fluent Bit with the new plugin

```sh
$ bin/fluent-bit -e /path/to/out_pulsar.so -i cpu -o pulsar-go
```

Configuration File

```sh
$ bin/fluent-bit -c /fluent-bit/etc/fluent-bit.conf
```

In addition download the following data sample file (130KB):

- https://fluentbit.io/samples/sp-samples-1k.log

```sh
$ curl -LO https://fluentbit.io/samples/sp-samples-1k.log
```

Ref: [Hands On! 101 - Fluent Bit: Official Manual](https://docs.fluentbit.io/manual/stream-processing/getting-started/hands-on)
