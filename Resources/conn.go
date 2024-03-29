package resources

import (
	"time"

	config "github.com/datrine/alumni_business/Config"
	models "github.com/datrine/alumni_business/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dbURL := config.GetDBURL()
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		}})
	db.AutoMigrate(&models.Account{},
		&models.Post{}, &models.Comment{}, &models.Transaction{})
	return db, err
}
