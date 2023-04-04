package impl

import (
	"backend/api/models"
	"backend/vendor/gorm.io/gorm"
)

type PetugasRepositoryImpl struct {
	DB *gorm.DB
}

// lihat data

func (r *PetugasRepositoryImpl) PetugasView() ([]models.Petugas, error) {
	row, err := r.DB.Raw("select * FROM petugas").Rows()

	if err != nil {
		return nil, err
	}
	var petugass []models.Petugas
	for row.Next() {
		var petugas models.Petugas
		err := row.Scan(
			&petugas.ID,
			&petugas.NamaPetugas,
			&petugas.AlamatPetugas,
			&petugas.JenisKelamin,
			&petugas.JamKerja,
			&petugas.JamPulang,
			&petugas.Gaji,
			&petugas.CreatedAt,
			&petugas.UpdatedAt)
		if err != nil {
			return nil, err
		}
		petugass = append(petugass, petugas)
	}
	return petugass, nil
}

func (r *PetugasRepositoryImpl) FindByName(NamaPetugas string) (models.Petugas, error) {
	var petugas models.Petugas
	r.DB.Model(petugas).Where("nama_petugas=?", NamaPetugas).Scan(&petugas)
	return petugas, nil
}
