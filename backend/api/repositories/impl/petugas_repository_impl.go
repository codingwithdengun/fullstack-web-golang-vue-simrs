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

// mencari data berdasarkan nama
func (r *PetugasRepositoryImpl) FindByName(NamaPetugas string) (models.Petugas, error) {
	var petugas models.Petugas
	r.DB.Model(petugas).Where("nama_petugas=?", NamaPetugas).Scan(&petugas)
	return petugas, nil
}

// mencari data berdasarkan id
func (r *PetugasRepositoryImpl) FindById(id uint) (models.Petugas, error) {
	var petugas models.Petugas
	r.DB.Model(petugas).Where("id=?", id).Scan(petugas)
	return petugas, nil
}

// menambakan data
func (r *PetugasRepositoryImpl) createPetugas(data models.PetugasRequest) (bool, error) {
	err := r.DB.Exec("INSERT INTO petugas (nama_petugas, alamat_petugas, jenis_kelamin, jam_kerja, jam_pulang, gaji) VALUES(?,?,?,?,?,?)", data.NamaPetugas, data.AlamatPetugas, data.JenisKelamin, data.JamKerja, data.JamPulang, data.Gaji).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// update data
func (r *PetugasRepositoryImpl) updatePetugas(data models.PetugasRequest, id uint) (bool, error) {
	err := r.DB.Exec("UPDATE petugas SET nama_petugas=?, alamat_petugas=?, jenis_kelamin=?,jam_kerja=?,jam_pulang=?,gaji=?", data.NamaPetugas, data.AlamatPetugas, data.JenisKelamin, data.JamKerja, data.JamPulang, data.Gaji, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// delete data
func (r *PetugasRepositoryImpl) deletePetugas(id uint) (bool, error) {
	err := r.DB.Where("id=?", id).Delete(&models.Petugas{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
