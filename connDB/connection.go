package connDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection opens the connection with the database
func Connection() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3008)/PadawanDB?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
