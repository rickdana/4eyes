package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rickdana/4eyes-poc/4eyes/model"
	"github.com/rickdana/4eyes-poc/4eyes/service"
)

type FourEyesHandler struct {
	fourEyesService service.FourEyesService
	kafkaClient     *service.KafkaClient
}

func NewFourEyesHandler(fourEyesService service.FourEyesService, kafkaClient *service.KafkaClient) *FourEyesHandler {
	return &FourEyesHandler{fourEyesService: fourEyesService, kafkaClient: kafkaClient}
}

func (f *FourEyesHandler) getAllFourEyes(c *fiber.Ctx) error {
	fourEyesReviews, err := f.fourEyesService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fourEyesReviews)
}

func (f *FourEyesHandler) getFourEyesById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	fourEyesReview, err := f.fourEyesService.Get(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("FourEyes with id %d not found", id),
			})
		}
	}
	return c.JSON(fourEyesReview)
}

func (f *FourEyesHandler) getFourEyesByBoType(c *fiber.Ctx) error {
	boType := c.Params("boType")

	params := map[string]string{
		"status":    c.Query("status"),
		"requester": c.Query("requester"),
		"reviewer":  c.Query("reviewer"),
	}

	fourEyesReviews, err := f.fourEyesService.GetAllByBoType(boType, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fourEyesReviews)
}

func (f *FourEyesHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	fourEyesReview, err := f.fourEyesService.Get(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("FourEyes with id %d not found", id),
			})
		}
	}

	var reviewStatus model.ReviewStatus
	err = c.BodyParser(&reviewStatus)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if reviewStatus.Action == "APPROVED" {
		fourEyesReview.Status = "ACTIVE"
	}

	if reviewStatus.Action == "REJECTED" {
		fourEyesReview.Status = "REJECTED"
	}

	f.kafkaClient.Send("fourEyes", fourEyesReview)

	updatedFourEyesReview, err := f.fourEyesService.Update(fourEyesReview)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(updatedFourEyesReview)
}
