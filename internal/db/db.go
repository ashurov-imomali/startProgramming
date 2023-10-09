package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/configs"
)

func GetConnection(configs *configs.Configs) (*gorm.DB, error) {
	dbSettings := "host=" + configs.DbHost + " port =" + configs.DbPort + " user=" + configs.UserName +
		" password=" + configs.Password + " dbname=" + configs.Database
	db, err := gorm.Open(postgres.Open(dbSettings), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
