package config



import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMySql(dsn string) (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return
}

