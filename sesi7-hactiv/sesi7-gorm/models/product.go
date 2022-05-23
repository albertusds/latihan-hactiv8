package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique;type:varchar(191)"`
	Brand     string `gorm:"not null;unique;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("Before insert to table users")
	if len(p.Name) < 2 {
		err = fmt.Errorf("your name is too short")
	}
	return
}

func (p *Product) Print() {
	fmt.Println("ID :", p.ID)
	fmt.Println("Name :", p.Name)
	fmt.Println("Brand :", p.Brand)
	fmt.Println("")
}
