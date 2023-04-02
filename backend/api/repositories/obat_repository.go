package repositories

import "backend/api/models"

type ObatRepository interface {
	ObatView() ([]models.Obat, error)
	FindById(uint) (models.Obat, error)
	FindByNama(string) (models.Obat, error)
	Create(models.ObatRequest) (bool, error)
	Update(models.ObatRequest, uint) (bool, error)
	Delete(uint) (bool, error)
}
