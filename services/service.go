package services

import (
	"photoApi/models"
	"photoApi/repository"
)

type DefualtService struct {
	Repo repository.PhotoRepsitory
}

type Service interface {
	PhotoGetAll() ([]models.Photo, error)
}

func (t DefualtService) PhotoGetAll() ([]models.Photo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewPhotoService(Repo repository.PhotoRepsitory)DefualtService  {
	return DefualtService{Repo: Repo}
}
