package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() error {
	dsn := fmt.Sprintf("host=%v user=%v port=%v password=%v dbname=%v sslmode=%v",
		Conf.Database.Host,
		Conf.Database.User,
		Conf.Database.Port,
		Conf.Database.Password,
		Conf.Database.Dbname,
		Conf.Database.Sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}
	Db = db
	return nil
}
