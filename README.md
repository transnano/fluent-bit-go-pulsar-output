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

## Playground

The output produced should resemble the following:

```sh
Fluent Bit v1.5.3
* Copyright (C) 2019-2020 The Fluent Bit Authors
* Copyright (C) 2015-2018 Treasure Data
* Fluent Bit is a CNCF sub-project under the umbrella of Fluentd
* https://fluentbit.io
...
time="2020-08-20T14:11:34Z" level=info msg="Connecting to broker" remote_addr="pulsar://pulsar:6650"
time="2020-08-20T14:11:34Z" level=info msg="TCP connection established" local_addr="172.20.0.3:36542" remote_addr="pulsar://pulsar:6650"
time="2020-08-20T14:11:34Z" level=info msg="Connection is ready" local_addr="172.20.0.3:36542" remote_addr="pulsar://pulsar:6650"
time="2020-08-20T14:11:34Z" level=info msg="Connecting to broker" remote_addr="pulsar://pulsar:6650"
time="2020-08-20T14:11:34Z" level=info msg="TCP connection established" local_addr="172.20.0.3:36544" remote_addr="pulsar://pulsar:6650"
time="2020-08-20T14:11:34Z" level=info msg="Connection is ready" local_addr="172.20.0.3:36544" remote_addr="pulsar://pulsar:6650"
time="2020-08-20T14:11:34Z" level=info msg="Created producer" cnx="172.20.0.3:36544 -> 172.20.0.2:6650" producer_name=standalone-0-4 topic="persistent://public/default/test"
2020/08/20 14:11:34 [pulsar-go][info][Init] Succeeded: pulsar://pulsar:6650, test
...
2020/08/20 14:11:35 [pulsar-go][debug][FlushCtx] JSON: {"ip": 73.113.230.135, "word": balsamine, "country": Japan, "flag": false, "num": 96, "date": 22/abr/2019:12:43:51 -0600, }
2020/08/20 14:11:35 [pulsar-go][debug][FlushCtx] JSON: {"ip": 85.61.182.212, "word": elicits, "country": Argentina, "flag": true, "num": 73, "date": 22/abr/2019:12:43:52 -0600, }
2020/08/20 14:11:35 [pulsar-go][debug][FlushCtx] JSON: {"date": 22/abr/2019:12:43:52 -0600, "ip": 18.135.244.142, "word": chesil, "country": Argentina, "flag": true, "num": 19, }
...
2020/08/20 14:11:46 [pulsar-go][info][FlushCtx] Succeeded: 1000
```
