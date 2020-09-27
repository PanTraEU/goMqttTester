package configUtil

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	UseAuth      bool
	MqttUser     string
	MqttPassword string
	MqttHost     string
	Topic        string
	UseTls       bool
}

func GetConfig(configfile string) Configuration {
	file, _ := os.Open(configfile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
