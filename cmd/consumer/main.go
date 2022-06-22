package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	//Create consumer
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err)
		return
	}
	//Set partition
	partitionList, err := consumer.Partitions("chat-room")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	//Cyclic partition
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("chat-room", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}

		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
