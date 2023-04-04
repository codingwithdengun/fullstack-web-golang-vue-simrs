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
func (r *PetugasServiceImpl) FindById(id uint) (models.Petugas, error) {
	get, err := r.petugasRepo.PetugasFinById(id)
	if err != nil {
		return models.Petugas{}, err
	}
	return get, err
}

// cari data berdasarkan nama
func (r *PetugasServiceImpl) FindByPetugas(name string) (models.Petugas, error) {
	get, err := r.petugasRepo.PetugasFindByName(name)

	if err != nil {
		return models.Petugas{}, err
	}
	return get, err
}

// proses update petugas
func (r *PetugasServiceImpl) UpdatePetugas(data models.PetugasRequest, id uint) (bool, error) {
	get, err := r.petugasRepo.PetugasUpdate(data, id)
	if err != nil {
		return false, err
	}
	return get, nil
}

// proses delete petugas
func (r *PetugasServiceImpl) DeletePetugas(id uint) (bool, error) {
	get, err := r.petugasRepo.PetugasDelete(id)
	if err != nil {
		return false, err
	}
	return get, err
}
