package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	//db menampung data db, err menampung data error
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_kp?parseTime=true"), &gorm.Config{}) //null variabel

	if err != nil {
		panic(err)
	}
	db.Exec("SET time_zone = 'SYSTEM'")
	// db.AutoMigrate(&Surat{})

	DB = db
}
