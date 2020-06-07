package migrations

import (
	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/database/seeders"
	"github.com/jinzhu/gorm"
)

/*Should contain all models*/
func allModels() *[]interface{} {
	a := []interface{}{
		&models.User{},
		//TODO add others for migration
	}
	return &a
}

/*DefaultDestructiveReset Drops all existing models and remigrates them*/
func DefaultDestructiveReset(db *gorm.DB) error {
	return DestructiveReset(db, *allModels()...)
}

/*DefaultMigration migrates all existing models*/
func DefaultMigration(db *gorm.DB) error {
	return AutoMigrate(db, *allModels()...)
}

/*DestructiveReset drop all tables and remigrate*/
func DestructiveReset(db *gorm.DB, models ...interface{}) error {
	err := db.DropTableIfExists(models...).Error
	if err != nil {
		return err
	}
	return AutoMigrate(db, models...)
}

/*AutoMigrate softly migrate given models*/
func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...).Error
	seeders.DefaultSeed()
	return err
}
