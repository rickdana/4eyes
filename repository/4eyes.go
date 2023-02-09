package repository

import (
	"github.com/rickdana/4eyes-poc/4eyes/model"
	"gorm.io/gorm"
	"time"
)

type FourEyesRepository interface {
	Saved(fourEyes model.FourEyesReview) (model.FourEyesReview, error)
	Find(id uint) (model.FourEyesReview, error)
	FindByBoIdAndBoType(boId uint, boType string) (model.FourEyesReview, error)
	FindAll() ([]model.FourEyesReview, error)
	FindAllByBoType(boType string, params map[string]string) ([]model.FourEyesReview, error)
	Update(fourEyes model.FourEyesReview) (model.FourEyesReview, error)
}

type FourEyesRepoImpl struct {
	db *gorm.DB
}

func NewFourEyesRepoImpl(db *gorm.DB) FourEyesRepository {
	return &FourEyesRepoImpl{db: db}
}

func (f *FourEyesRepoImpl) Saved(fourEyes model.FourEyesReview) (model.FourEyesReview, error) {
	result := f.db.Create(&fourEyes)
	return fourEyes, result.Error
}

func (f *FourEyesRepoImpl) Find(id uint) (model.FourEyesReview, error) {
	var fourEyes model.FourEyesReview
	result := f.db.First(&fourEyes, id)
	return fourEyes, result.Error
}

func (f *FourEyesRepoImpl) FindByBoIdAndBoType(boId uint, boType string) (model.FourEyesReview, error) {
	var fourEyes model.FourEyesReview
	result := f.db.Where("bo_id = ? AND bo_type = ?", boId, boType).First(&fourEyes)
	return fourEyes, result.Error
}

func (f *FourEyesRepoImpl) FindAll() ([]model.FourEyesReview, error) {
	var fourEyes []model.FourEyesReview
	result := f.db.Find(&fourEyes)
	return fourEyes, result.Error
}

func (f *FourEyesRepoImpl) FindAllByBoType(boType string, params map[string]string) ([]model.FourEyesReview, error) {
	var fourEyes []model.FourEyesReview
	params["bo_type"] = boType
	result := f.db.Where(params).Find(&fourEyes)
	return fourEyes, result.Error
}

func (f *FourEyesRepoImpl) Update(fourEyes model.FourEyesReview) (model.FourEyesReview, error) {
	now := time.Now()
	f.db.Raw("UPDATE four_eyes SET status = ?, updated_at = ? WHERE id = ?", fourEyes.Status, now, fourEyes.ID)
	return fourEyes, nil
}
