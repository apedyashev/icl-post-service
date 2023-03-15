package datastore

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// DBMS := "mysql"

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", DbHost, DbPort, DbUser, DbPassword, DbName)
	fmt.Println("DSN:", dsn)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// mySqlConfig := &mysql.Config{
	// 	User:                 config.C.Database.User,
	// 	Passwd:               config.C.Database.Password,
	// 	Net:                  config.C.Database.Net,
	// 	Addr:                 config.C.Database.Addr,
	// 	DBName:               config.C.Database.DBName,
	// 	AllowNativePasswords: config.C.Database.AllowNativePasswords,
	// 	Params: map[string]string{
	// 		"parseTime": config.C.Database.Params.ParseTime,
	// 	},
	// }

	// db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
