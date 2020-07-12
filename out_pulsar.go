package main

import "C"
import (
	"context"
	"fmt"
	"unsafe"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/fluent/fluent-bit-go/output"
)

type pulsarConfig struct {
	BrokerServiceUrl           string
	Tennant                    string
	Namespace                  string
	Topic                      string
	TlsEnabled                 bool
	TlsTrustCertsFilePath      string
	TlsAllowInsecureConnection bool
}

type pulsarClient struct {
	Client   *pulsar.Client
	Producer *pulsar.Producer
}

var client *pulsarClient

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded
	return output.FLBPluginRegister(ctx, "flb-go-pulsar", "Output to Apache pulsar")
}

//export FLBPluginInit
func FLBPluginInit(plugin unsafe.Pointer) int {

	config := &pulsarConfig{
		BrokerServiceUrl:           output.FLBPluginConfigKey(plugin, "BrokerServiceUrl"),
		Tennant:                    output.FLBPluginConfigKey(plugin, "Tennant"),
		Namespace:                  output.FLBPluginConfigKey(plugin, "Namespace"),
		Topic:                      output.FLBPluginConfigKey(plugin, "Topic"),
		TlsEnabled:                 parseBool(output.FLBPluginConfigKey(plugin, "TlsEnabled")),
		TlsTrustCertsFilePath:      output.FLBPluginConfigKey(plugin, "TlsTrustCertsFilePath"),
		TlsAllowInsecureConnection: parseBool(output.FLBPluginConfigKey(plugin, "TlsAllowInsecureConnection")),
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: config.BrokerServiceUrl,
	})
	if err != nil {
		return output.FLB_ERROR
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	if err != nil {
		return output.FLB_ERROR
	}

	client = &pulsarClient{
		Client:   client,
		Producer: producer,
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
		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(record),
		})

		if err != nil {
			fmt.Printf("[flb-go::pulsar][error] err: %s\n", err)
			return output.FLB_ERROR
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
