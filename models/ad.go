package models

//Ad ...
type Ad struct {
	ID        uint     `json:"id" gorm:"primary_key"`
	Name      string   `json:"name"`
	Descript  string   `json:"desc"`
	Price     float64  `json:"price"`
	CreatedAt string   `json:"created_at"`
	MainPic   string   `json:"mainpic"`
	OtherPics []string `json:"pic_links"`
}
