# fluent-bit-go-pulsar-output-sample

Fluent-Bit go apache pulsar output plugin

## Build a Go Plugin

To build the code above, use the following line:

```sh
$ go build -buildmode=c-shared -o out_pulsar.so out_pulsar.go
```

Once built, a shared library called `out_pulsar.so` will be available. It's really important to double check the final .so file is what we expect. Doing a ldd over the library we should see something similar to this:

```sh

```

## Run Fluent Bit with the new plugin

```sh
$ bin/fluent-bit -e /path/to/out_gstdout.so -i cpu -o gstdout
```
