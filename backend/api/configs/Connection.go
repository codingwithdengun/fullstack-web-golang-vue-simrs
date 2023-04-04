package configs

import (
	"backend/api/models"
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	val := url.Values{}

	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", Config.Database.DBHost, Config.Database.DBPort, Config.Database.DBUser, Config.Database.DBName, Config.Database.DBPass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot Connected Database", err.Error())
		return nil
	}
	err = db.AutoMigrate(
		&models.Petugas{},
		&models.DetailFaktur{},
		&models.DetailRekamMedis{},
		&models.Dokter{},
		&models.FakturPembayaran{},
		&models.KamarInap{},
		&models.Obat{},
		&models.Pasien{},
		&models.Pendaftaran{},
		&models.Perawat{},
		&models.Poli{},
		&models.RekamMedis{},
	)

	if err != nil {
		log.Fatal("Failed Migrated ")
		return nil
	}
	sqlDB, err := db.DB()
	err = sqlDB.Ping()

	if err != nil {
		log.Fatal("Request Timeout ", err)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	log.Info("Connected Database" + Config.Database.DBDriver)
	return db
}
