package model

type User struct {
	ID       int    `gorm:"primarykey;auto_increment"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Email    string `json:"email"    gorm:"unique"`
}
