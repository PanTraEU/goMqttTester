package main

import (
	"../configUtil"
	"../tlsutils"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./config.json", "Specify config file. Default is 'config.json'.")
	flag.Parse()

	conf := configUtil.GetConfig(configFile)

	mqttHost := conf.MqttHost
	producerTopic := conf.Topic
	prodId := "test_producer_1"

	optsProd := TlsUtils.MqttOpts(mqttHost, prodId, conf, true)

	prodClient := mqtt.NewClient(optsProd)
	if token := prodClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var counter uint64 = 0
	for true {
		text := fmt.Sprintf("%d", counter)
		token := prodClient.Publish(producerTopic, 2, false, text)
		token.Wait()

		log.Println(fmt.Sprintf("send new daykey: %s", text))
		time.Sleep(2 * time.Second)
		counter++
	}

	defer prodClient.Disconnect(250)
}
