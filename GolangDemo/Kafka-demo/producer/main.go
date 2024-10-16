package main

import (
	"crypto/rand"
	"log"

	"github.com/IBM/sarama"
)

var (
	topic = "test-topic"
)

func main() {
	// cmd := exec.Command("C:/logstash-8.15.2/bin/logstash.bat", "-f", "C:/Users/admin/Desktop/Golang/worker-crawl-cpe/config/logstash.conf")
	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatalf("Failed to start Logstash: %v", err)
	// }
	p:= make([]byte, 10)
	rand.Read(p)
	pushMessagetoTopic(topic, p)
}

func connectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func pushMessagetoTopic(topic string, message []byte) error {
	brokers := []string{"localhost:9093"}
	producer, err := connectProducer(brokers)
	if err != nil {
		return err
	}
	defer producer.Close()

	//create a new message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	//send message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("stored in topic %s / partition %d /offset %d \n", topic, partition, offset)
	return nil
}
