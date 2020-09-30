#!/usr/bin/env sh

echo "-----------------------------------------------------"
echo "init consumer"
echo "-----------------------------------------------------"
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con1 -m=init
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con2 -m=init
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con3 -m=init
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con4 -m=init
echo "-----------------------------------------------------"

#sleep 5
#
#echo "-----------------------------------------------------"
#echo "run producer"
#echo "-----------------------------------------------------"
#go run mqttProducer/main.go -c="./configs/tlsProducerConfig.json"
#echo "-----------------------------------------------------"

sleep 5

echo "-----------------------------------------------------"
echo "run consumer"
echo "-----------------------------------------------------"
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con1 -m=consume
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con2 -m=consume
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con3 -m=consume
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con4 -m=consume
echo "-----------------------------------------------------"

sleep 15

echo "-----------------------------------------------------"
echo "run consumer"
echo "-----------------------------------------------------"
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con1 -m=consume
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con2 -m=consume
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con3 -m=consume
go run mqttConsumer/main.go -c="./configs/tlsConsumerConfig.json" -cid=con4 -m=consume
echo "-----------------------------------------------------"
