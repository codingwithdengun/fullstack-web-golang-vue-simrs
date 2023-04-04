package models

import "time"

type Petugas struct {
	ID            uint      `json:"id" gorm:"column:id; size:36;" sql:"type:uuid;primaryKey:uuid_generate_v4()"`
	NamaPetugas   string    `gorm:"not null"`
	AlamatPetugas string    `gorm:"not null"`
	JenisKelamin  string    `gorm:"not null"`
	JamKerja      time.Time `gorm:"not null"`
	JamPulang     time.Time `gorm:"not null"`
	Gaji          float32   `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

type PetugasRequest struct {
	NamaPetugas   string    `json:"nama_petugas" binding:"required"`
	AlamatPetugas string    `json:"alamat_petugas" binding:"required"`
	JenisKelamin  string    `json:"jenis_kelamin" binding:"required"`
	JamKerja      time.Time `json:"jam_kerja" binding:"required"`
	JamPulang     time.Time `json:"jam_pulang" binding:"required"`
	Gaji          float32   `json:"gaji" binding:"required"`
}

type PetugasResponse struct {
	ID            uint      `json:"id"`
	NamaPetugas   string    `json:"nama_petugas"`
	AlamatPetugas string    `json:"alamat_petugas"`
	JenisKelamin  string    `json:"jenis_kelamin"`
	JamKerja      time.Time `json:"jam_kerja"`
	JamPulang     time.Time `json:"jam_pulang"`
	Gaji          float32   `json:"gaji"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
