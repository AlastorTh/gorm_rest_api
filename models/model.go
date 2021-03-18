package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Ad ...
type Ad struct {
	gorm.Model
	Name      string   `json:"name"`
	Descript  string   `json:"desc"`
	Price     float64  `json:"price"`
	MainPic   string   `json:"mainpic"`
	OtherPics []string `gorm:"type:varchar(64)[]" json:"pic_links"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
