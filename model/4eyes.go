package model

import "gorm.io/gorm"

var (
	InitialStatus = map[string]string{
		"CREATE": "pendingCreation",
		"UPDATE": "pendingDeletion",
		"DELETE": "pendingModification",
	}
)

type FourEyesStatus struct {
	Approved string `json:"approved"`
	Rejected string `json:"rejected"`
}

type FourEyesFlow map[string]FourEyesStatus

type FourEyesReview struct {
	gorm.Model
	Status        string `json:"status"`
	BoId          uint   `json:"boId"`
	BoType        string `json:"boType"`
	Before        string `json:"before"`
	After         string `json:"after"`
	Reviewer      string `json:"reviewer"`
	Requester     string `json:"requester"`
	ReviewComment string `json:"reviewComment"`
}

func NewFourEyesReview(boId uint, boType string, before string, after string) *FourEyesReview {
	return &FourEyesReview{BoId: boId, BoType: boType, Before: before, After: after}
}

type ReviewStatus struct {
	Action  string `json:"action"`
	Comment string `json:"comment"`
}
