package main

import (
	"../configUtil"
	"../tlsutils"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/utils"
	"log"
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

	// sending 10 messages
	for i := 0; i < 5; i++ {
		text := string(utils.UUID())
		token := prodClient.Publish(producerTopic, 2, false, text)
		token.Wait()
		log.Println(fmt.Sprintf("send new daykey: %s", text))
	}
	prodClient.Disconnect(250)
}
