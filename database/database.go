package database

import (
	"fmt"
	"log"
	"todolist/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectionDatabase() (*gorm.DB, error) {

	viper.SetConfigFile("./env/developing.env")
	viper.ReadInConfig()

	//Database koneksi PostgreSql
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	address := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	Db, err = gorm.Open(postgres.Open(address), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	Db.AutoMigrate(&models.List{}, &models.SubList{})
	return Db, nil
}
