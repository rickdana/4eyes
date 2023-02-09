package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type sender interface {
	Send(key string, message any) error
}

type receiver interface {
	Receive() ([]any, error)
}

type KafkaConfig struct {
	Url       string
	Topic     string
	Partition int
}

type KafkaClient struct {
	Conn *kafka.Conn
}

func NewKafkaClient(config KafkaConfig) *KafkaClient {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.Url, config.Topic, config.Partition)
	if err != nil {
		return nil
	}
	return &KafkaClient{Conn: conn}
}

func (k *KafkaClient) Send(key string, message any) error {
	//TODO implement me
	panic("implement me")
}

func (k *KafkaClient) Receive() ([]any, error) {
	k.Conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := k.Conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	defer batch.Close()
	defer k.Conn.Close()

	for {
		message, err := batch.ReadMessage()

		if err != nil {
			log.Println("Error reading message from kafka", err)
		}
		fmt.Println("Received message", string(message.Value))
		fmt.Println("Message key", string(message.Key))
		fmt.Println("Message value", string(message.Value))
	}
}
