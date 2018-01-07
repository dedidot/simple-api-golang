package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	Codes int `gorm:"column:codes"`
	Name string `gorm:"column:name" json:"name"`
	Author string   `gorm:"column:author" json:"author"`
	Category string `gorm:"column:category" json:"category"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{})
	hasUser := db.HasTable(&Book{})
	fmt.Println("Tabel book is ", hasUser)
	if !hasUser {
		db.CreateTable(&Book{})
	}

	return db
}
