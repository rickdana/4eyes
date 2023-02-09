package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rickdana/4eyes-poc/4eyes/model"
	"github.com/rickdana/4eyes-poc/4eyes/repository"
	"github.com/rickdana/4eyes-poc/4eyes/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func intDB() (*gorm.DB, error) {
	dsn := "host=localhost user=dbuser password=4eyesP@ssw0rd dbname=4eyes_receiver port=5432 sslmode=disable TimeZone=Europe/Paris"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),

	})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(model.FourEyesReview{})
	return db, err
}

func main() {
	db, err := intDB()

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

	if err != nil {
		panic("failed to connect database")
	}

	fourEyesHandler := NewFourEyesHandler(fourEyesService, kafkaClient)
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/four-eyes", fourEyesHandler.getAllFourEyes)

	v1.Get("/four-eyes/:id", fourEyesHandler.getFourEyesById)
	v1.Get("/four-eyes/:boType", fourEyesHandler.getFourEyesByBoType)

	v1.Post("/four-eyes/:id", fourEyesHandler.Update)

	log.Fatal(app.Listen(":3010"))
}
