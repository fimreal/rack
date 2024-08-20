package models

import (
	"github.com/fimreal/rack/pkg/components/db"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ORM struct {
	*gorm.DB
}

func NewORM() (*ORM, error) {
	driver := viper.GetString("db_driver")
	host := viper.GetString("db_host")
	port := viper.GetInt("db_port")
	user := viper.GetString("db_user")
	password := viper.GetString("db_password")
	dbname := viper.GetString("db_name")

	orm := db.NewDatabaseConfig(driver, user, password, host, port, dbname)
	i, err := db.NewInstance(orm)
	return &ORM{DB: i}, err
}
