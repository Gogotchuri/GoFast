package database

import (
	"fmt"
	"github.com/Gogotchuri/GoFast/config"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


/*Database access object*/
var dao *gorm.DB = nil
var once sync.Once

/*GetInstance Returns singleton database*/
func GetInstance() *gorm.DB {
	once.Do(func() {
		dao = initDB(config.GetInstance())
	})
	return dao
}

/*Establishes connection to database*/
func initDB(cfg *config.Config) *gorm.DB {
	if cfg == nil {
		panic(fmt.Errorf("database config not set"))
	}
	db, err := gorm.Open(cfg.Database.Dialect(), cfg.Database.DBConnectionInfo())
	if err != nil {
		panic(err)
	}
	return db
}
