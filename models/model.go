package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

//Ad ...
type Ad struct {
	gorm.Model
	Name      string         `json:"name"`
	Descript  string         `json:"desc"`
	Price     float64        `json:"price"`
	MainPic   string         `json:"mainpic"`
	OtherPics pq.StringArray `gorm:"type:varchar(64)[]" json:"other_pics"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type CreateAdInput struct {
	Name      string         `json:"name"`
	Descript  string         `json:"desc"`
	Price     float64        `json:"price"`
	MainPic   string         `json:"mainpic"`
	OtherPics pq.StringArray `gorm:"type:varchar(64)[]" json:"pic_links"`
}
