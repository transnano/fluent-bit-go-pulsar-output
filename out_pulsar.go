package main

import "C"
import (
	"context"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/fluent/fluent-bit-go/output"
)

type pulsarClient struct {
	Client   pulsar.Client
	Producer pulsar.Producer
}

var client *pulsarClient

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded
	return output.FLBPluginRegister(ctx, "flb-go-pulsar", "Output to Apache pulsar")
}

//export FLBPluginInit
func FLBPluginInit(plugin unsafe.Pointer) int {
	clientOpts := pulsar.ClientOptions{
		URL: output.FLBPluginConfigKey(plugin, "BrokerServiceUrl"),
		// TLSTrustCertsFilePath:      output.FLBPluginConfigKey(plugin, "TLSTrustCertsFilePath"),
		// TLSAllowInsecureConnection: parseBool(output.FLBPluginConfigKey(plugin, "TLSAllowInsecureConnection")),
	}
	pClient, err := pulsar.NewClient(clientOpts)
	if err != nil {
		fmt.Errorf("[flb-go-pulsar][error][Init] failed: %s, %v\n", clientOpts.URL, err)
		return output.FLB_ERROR
	}
	producerOpts := pulsar.ProducerOptions{
		Topic: output.FLBPluginConfigKey(plugin, "Topic"),
		// CompressionType: pulsar.LZ4,
	}
	pProducer, err := pClient.CreateProducer(producerOpts)
	if err != nil {
		fmt.Errorf("[flb-go-pulsar][error][Init] failed: %s, %v\n", producerOpts.Topic, err)
		return output.FLB_ERROR
	}
	fmt.Printf("[flb-go-pulsar][info][Init] Succeeded: %s, %s\n", clientOpts.URL, producerOpts.Topic)

	client = &pulsarClient{
		Client:   pClient,
		Producer: pProducer,
	}

	// Gets called only once for each instance you have configured.
	return output.FLB_OK
}

//export FLBPluginFlushCtx
func FLBPluginFlushCtx(ctx, data unsafe.Pointer, length C.int, tag *C.char) int {

	dec := output.NewDecoder(data, int(length))
	for {
		ret, _, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		for _, v := range record {
			var payload []byte
			switch t := v.(type) {
			case string:
				payload = []byte(t)
			case []byte:
				payload = t
			default:
				payload = []byte(fmt.Sprintf("%v", v))
			}
			fmt.Printf("[flb-go-pulsar][info][FlushCtx] presend: %s\n", string(payload))

			_, err := client.Producer.Send(context.Background(), &pulsar.ProducerMessage{
				Payload: payload,
			})
			if err != nil {
				fmt.Errorf("[flb-go-pulsar][error][FlushCtx] err: %v\n", err)
				return output.FLB_ERROR
			}
		}
	}

	// Gets called with a batch of records to be written to an instance.
	return output.FLB_OK
}

func parseBool(s string) bool {
	toBool, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return toBool
}

//export FLBPluginExit
func FLBPluginExit() int {
	if client != nil {
		if client.Producer != nil {
			client.Producer.Close()
		}
		if client.Client != nil {
			client.Client.Close()
		}
	}
	return output.FLB_OK
}

func main() {
}
