package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

/*
	Initialize NewConfig configuration sarama.NewConfig
	Create producer sarama.NewSyncProducer
	Create message sarama.ProducerMessage
	Send message producer.SendMessage
*/
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "chat-room"
	msg.Key = sarama.StringEncoder("Key")
	msg.Value = sarama.StringEncoder("Hello")

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer producer.Close()

	for i := 1; i <= 10; i++ {
		pid, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			panic(err)
		}
		fmt.Printf("pid: %v, offset: %v, topic: %v, key: %v, msg: %v\n", pid, offset, msg.Topic, msg.Key, msg.Value)
	}

}
