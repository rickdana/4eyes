package graph

import "github.com/rickdana/4eyes-poc/4eyes/service"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	FourEyesService service.FourEyesService
	EventService    service.KafkaClient
}
