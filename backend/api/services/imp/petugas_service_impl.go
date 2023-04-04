package imp

import (
	"backend/api/models"
	"backend/api/repositories"
	"backend/api/services"
)

type PetugasServiceImpl struct {
	petugasRepo repositories.PetugasRepository
}

func CreatePetugasService(petugasRepo repositories.PetugasRepository) services.PetugasService {
	return &PetugasServiceImpl{
		petugasRepo: petugasRepo,
	}
}

// proses lihat petugas
func (r *PetugasServiceImpl) ViewUser() ([]models.Petugas, error) {
	get, err := r.petugasRepo.PetugasView()
	if err != nil {
		return nil, err
	}
	return get, nil
}

// cari data by id
func (r *PetugasServiceImpl) FindById(id uint) (models.Petugas, error)

// cari data berdasarkan nama
func (r *PetugasServiceImpl) FindByPetugas(name string) (models.Petugas, error)

// proses update petugas
func (r *PetugasServiceImpl) UpdatePetugas(data models.PetugasRequest, id uint) (bool, error)

// proses delete petugas
func (r *PetugasServiceImpl) DeletePetugas(id uint) (bool, error)
