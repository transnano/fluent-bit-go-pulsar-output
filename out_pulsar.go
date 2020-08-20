package main

import "C"
import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/fluent/fluent-bit-go/output"
)

var defaultMap = map[string]string{
	"BrokerServiceUrl":           "pulsar://pulsar:6650",
	"Tennant":                    "pulsar",
	"Namespace":                  "default",
	"Topic":                      "test",
	"TLSEnabled":                 "true",
	"TLSTrustCertsFilePath":      "file://path/to/cert",
	"TLSAllowInsecureConnection": "true",
}

type pulsarClient struct {
	Client   pulsar.Client
	Producer pulsar.Producer
}

var client *pulsarClient

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded
	return output.FLBPluginRegister(ctx, "pulsar-go", "Output to Apache pulsar")
}

//export FLBPluginInit
func FLBPluginInit(plugin unsafe.Pointer) int {
	outURL := getConfigKey(plugin, "BrokerServiceUrl")
	log.Printf("[pulsar-go][debug][Init] URL: %s\n", outURL)
	tennant := getConfigKey(plugin, "Tennant")
	log.Printf("[pulsar-go][debug][Init] Tennant: %s\n", tennant)
	clientOpts := pulsar.ClientOptions{
		URL:               outURL,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
		// TLSTrustCertsFilePath:      getConfigKey(plugin, "TLSTrustCertsFilePath"),
		// TLSAllowInsecureConnection: parseBool(getConfigKey(plugin, "TLSAllowInsecureConnection")),
	}
	pClient, err := pulsar.NewClient(clientOpts)
	if err != nil {
		log.Printf("[pulsar-go][error][Init] failed: %s, %v\n", clientOpts.URL, err)
		return output.FLB_ERROR
	}
	producerOpts := pulsar.ProducerOptions{
		Topic:           getConfigKey(plugin, "Topic"),
		CompressionType: pulsar.LZ4,
	}
	pProducer, err := pClient.CreateProducer(producerOpts)
	if err != nil {
		log.Printf("[pulsar-go][error][Init] failed: %s, %v\n", producerOpts.Topic, err)
		return output.FLB_ERROR
	}
	log.Printf("[pulsar-go][info][Init] Succeeded: %s, %s\n", clientOpts.URL, producerOpts.Topic)

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

	count := 0
	for {
		ret, _, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		b := &strings.Builder{}
		b.WriteString("{")

		for k, v := range record {
			var payload []byte
			switch t := v.(type) {
			case string:
				payload = []byte(t)
			case []byte:
				payload = t
			default:
				payload = []byte(fmt.Sprintf("%v", v))
			}
			// log.Printf("[pulsar-go][debug][FlushCtx] key: %v, value: %s\n", k, string(payload))

			s := fmt.Sprintf("\"%s\": %v, ", k, string(payload))
			b.WriteString(s)
		}
		b.WriteString("}")
		s := b.String()
		// log.Printf("[pulsar-go][debug][FlushCtx] JSON: %s\n", s)

		count++
		_, err := client.Producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(s),
		})
		if err != nil {
			log.Printf("[pulsar-go][error][FlushCtx] err: %v\n", err)
			return output.FLB_ERROR
		}
	}
	log.Printf("[pulsar-go][info][FlushCtx] Succeeded: %d\n", count)

	// Gets called with a batch of records to be written to an instance.
	return output.FLB_OK
}

func getConfigKey(plugin unsafe.Pointer, key string) string {
	s := output.FLBPluginConfigKey(plugin, key)
	if len(s) == 0 {
		// log.Printf("[pulsar-go][debug][getConfigKey] Default value")
		return defaultMap[key]
	}
	// log.Printf("[pulsar-go][debug][getConfigKey] Config value")
	return s
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
	log.Printf("[pulsar-go][info][exit] Graceful shutdown...")
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
