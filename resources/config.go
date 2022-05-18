package resources

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "admin:Team@1000@tcp(34.93.40.185:3306)/rpay"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func init() {
	Connect()
}
