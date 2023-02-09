package main

import (
	"fmt"
	"github.com/rickdana/4eyes-poc/4eyes/model"
	"github.com/rickdana/4eyes-poc/4eyes/repository"
	"github.com/rickdana/4eyes-poc/4eyes/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "host=localhost user=dbuser password=4eyesP@ssw0rd dbname=4eyes_receiver port=5432 sslmode=disable TimeZone=Europe/Paris"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),

	})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(model.FourEyesReview{})

	if err != nil {
		panic(err)
	}

	//initialize Repo
	repo := repository.NewFourEyesRepoImpl(db)

	//initialize service
	fourEyesService := service.NewFourEyesService(repo)

	kafkaConfig := service.KafkaConfig{
		Url:       "localhost:9092",
		Topic:     "create",
		Partition: 0,
	}
	kafkaClient := service.NewKafkaClient(kafkaConfig)

	batch := kafkaClient.Conn.ReadBatch(1, 57671680) // fetch 10KB min, 1MB max
	// fetch 10KB min, 1MB max

	run := true

	for run {
		message, err := batch.ReadMessage()

		if err != nil {
			if err.Error() == "EOF" {
				continue
			}
			log.Println("Error reading message from kafka", err)
			continue
		}
		fmt.Println("Received message", string(message.Value))
		fmt.Println("Message key", string(message.Key))
		fmt.Println("Message value", string(message.Value))

	}

	batch.Close()
	kafkaClient.Conn.Close()
}
