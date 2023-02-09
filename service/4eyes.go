package service

import (
	"github.com/rickdana/4eyes-poc/4eyes/model"
	"github.com/rickdana/4eyes-poc/4eyes/repository"
)

type FourEyesService interface {
	Saved(fourEyes model.FourEyesReview) (model.FourEyesReview, error)
	Get(id uint) (model.FourEyesReview, error)
	GetByBoIdAndBoType(boId uint, boType string) (model.FourEyesReview, error)
	GetAll() ([]model.FourEyesReview, error)
	GetAllByBoType(boType string, params map[string]string) ([]model.FourEyesReview, error)
	Update(fourEyes model.FourEyesReview) (model.FourEyesReview, error)
}

type FourEyesServiceImpl struct {
	fourEyesRepo repository.FourEyesRepository
}

func NewFourEyesService(fourEyesRepo repository.FourEyesRepository) FourEyesService {
	return &FourEyesServiceImpl{fourEyesRepo: fourEyesRepo}
}

func (f *FourEyesServiceImpl) Saved(fourEyes model.FourEyesReview) (model.FourEyesReview, error) {
	return f.fourEyesRepo.Saved(fourEyes)
}

func (f *FourEyesServiceImpl) Get(id uint) (model.FourEyesReview, error) {
	return f.fourEyesRepo.Find(id)
}

func (f *FourEyesServiceImpl) GetByBoIdAndBoType(boId uint, boType string) (model.FourEyesReview, error) {
	return f.fourEyesRepo.FindByBoIdAndBoType(boId, boType)
}

func (f *FourEyesServiceImpl) GetAll() ([]model.FourEyesReview, error) {
	return f.fourEyesRepo.FindAll()
}

func (f *FourEyesServiceImpl) GetAllByBoType(boType string, params map[string]string) ([]model.FourEyesReview, error) {
	return f.fourEyesRepo.FindAllByBoType(boType, params)
}

func (f *FourEyesServiceImpl) Update(fourEyes model.FourEyesReview) (model.FourEyesReview, error) {
	return f.fourEyesRepo.Update(fourEyes)
}
