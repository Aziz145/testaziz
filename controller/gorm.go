package controller

import (
	"github.com/Aziz145/testaziz/infra"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func ConnectDB() *InDB {
	// infra.LoadPostgreSQLDB()
	db := infra.LoadMySQLDB()
	// db := infra.LoadSQLiteDB()
	inDB := &InDB{DB: db}
	return inDB
}
