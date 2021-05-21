package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
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

func (a *Ad) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validation.Required, validation.Length(1, 200)),
		validation.Field(&a.Descript, validation.Required, validation.Length(1, 1000)),
		validation.Field(&a.Price, validation.Min(0.00)),
		validation.Field(&a.MainPic, validation.Required, validation.Length(1, 2000)),
		validation.Field(&a.OtherPics, validation.Required, validation.Length(1, 2), validation.Each(validation.Length(1, 2000))),
	)
}
