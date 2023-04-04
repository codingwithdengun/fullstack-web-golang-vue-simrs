package repositories

import "backend/api/models"

type PetugasRepository interface {
	PetugasView() ([]models.Petugas, error)
	PetugasFinById(uint) (models.Petugas, error)
	PetugasFindByName(string) (models.Petugas, error)
	PetugasCreate(models.PetugasRequest) (bool, error)
	PetugasUpdate(models.PetugasRequest, uint) (bool, error)
	PetugasDelete(uint) (bool, error)
}
