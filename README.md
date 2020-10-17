# fluent-bit-go pulsar output ![Releases](https://github.com/transnano/fluent-bit-go-pulsar-output/workflows/Releases/badge.svg) ![Publish Docker image](https://github.com/transnano/fluent-bit-go-pulsar-output/workflows/Publish%20Docker%20image/badge.svg) ![Vulnerability Scan](https://github.com/transnano/fluent-bit-go-pulsar-output/workflows/Vulnerability%20Scan/badge.svg) ![Haskell Dockerfile Linter](https://github.com/transnano/fluent-bit-go-pulsar-output/workflows/Haskell%20Dockerfile%20Linter/badge.svg)

![License](https://img.shields.io/github/license/transnano/fluent-bit-go-pulsar-output?style=flat)

![Container image version](https://img.shields.io/docker/v/transnano/fluent-bit-go-pulsar-output/latest?style=flat)
![Container image size](https://img.shields.io/docker/image-size/transnano/fluent-bit-go-pulsar-output/latest?style=flat)
![Container image pulls](https://img.shields.io/docker/pulls/transnano/fluent-bit-go-pulsar-output?style=flat)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/transnano/fluent-bit-go-pulsar-output)
[![Go Report Card](https://goreportcard.com/badge/github.com/transnano/fluent-bit-go-pulsar-output)](https://goreportcard.com/report/github.com/transnano/fluent-bit-go-pulsar-output)

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
2020/08/21 06:53:29 [pulsar-go][info][Register] version: v0.0.1, commit: 0d82f86
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

## Features

- [X] Non-TLS Connection
- [X] Persistent Topics
- [ ] Topic Compaction(Compression)
- [ ] Multi Tenancy
- [ ] Non-Persistent Topics
- [ ] Transport Encryption using TLS
- [ ] TLS Authentication
- [ ] Athenz (Authentication)
- [ ] Kerberos (Authentication)
- [ ] JSON Web Token Authentication

## Configuration Parameters

The plugin supports the following configuration parameters:

Key                        | Description                                                                        | Default
-------------------------- | ---------------------------------------------------------------------------------- | -------------------------
BrokerServiceUrl           | The brokers in the cluster you send data to                                        | `pulsar://localhost:6650`
Tennant                    | The topic's tenant within the instance.                                            | public
Namespace                  | Pulsar namespaces are logical groupings of topics.                                 | default
Topic                      | Topic which is a logical endpoint for publishing messages.                         | test
CompressionType            | The compression type of the message published by the producer                      | none
TLSEnabled                 | Encrypt communication with Apache Pulsar service                                   | false
TLSTrustCertsFilePath      | specify the path the trust cert file                                               | (none)
TLSAllowInsecureConnection | The client to connect to servers whose cert has not been signed by an approved CA. | false

Ref: [Messaging Concepts · Apache Pulsar](https://pulsar.apache.org/docs/en/2.6.0/concepts-messaging/#topics)
