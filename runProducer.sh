#!/usr/bin/env sh

echo "-----------------------------------------------------"
echo "run producer"
echo "-----------------------------------------------------"
go run mqttProducer/main.go -c="./configs/tlsProducerConfig.json"
echo "-----------------------------------------------------"
