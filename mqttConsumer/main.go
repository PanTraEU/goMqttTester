package main

import (
	"../configUtil"
	"../tlsutils"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var newMsgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	reader := client.OptionsReader()
	clientId := reader.ClientID()
	fmt.Printf("client_id: %s\t\ttopic: %s\t\t\"message: %s\t\tqos: %d\n", clientId, msg.Topic(), msg.Payload(), msg.Qos())
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "config.json", "Specify config file. Default is 'config.json'.")
	var consumerId string
	flag.StringVar(&consumerId, "cid", "test_consumer_1", "Specify consumerId. Default is 'test_consumer_1'.")
	var mode string
	flag.StringVar(&mode, "m", "init", "Specify mode [init|consume]. Default is 'init'.")
	flag.Parse()

	conf := configUtil.GetConfig(configFile)

	mqttHost := conf.MqttHost
	producerTopic := conf.Topic
	conId := consumerId

	optsConsumer := TlsUtils.MqttOpts(mqttHost, conId, conf, false)

	consumerClient := mqtt.NewClient(optsConsumer)
	if token := consumerClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if mode != "init" {
		fmt.Printf("consumer clientId : %s\t\ttopic: %s\n", conId, producerTopic)
		consumerClient.Subscribe(producerTopic, 2, newMsgHandler)
	} else {
		fmt.Printf("init clientId: %s\n", conId)

	}

	fmt.Printf("disconnecting\n")
	consumerClient.Disconnect(250)
}
