package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	DB, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
