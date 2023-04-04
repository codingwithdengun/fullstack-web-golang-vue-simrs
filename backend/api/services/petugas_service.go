package services

import "backend/api/models"

type PetugasService interface {
	ViewUser() ([]models.Petugas, error)
	FindById(uint) (models.Petugas, error)
	FindByPetugas(string) (models.Petugas, error)
	UpdatePetugas(models.PetugasRequest, uint) (bool, error)
	DeletePetugas(uint) (bool, error)
}
