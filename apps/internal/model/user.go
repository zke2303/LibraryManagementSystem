package user

import (
	"gorm.io/gorm"
)

type Gender uint8

const (
	GenderUnknown Gender = 0
	GenderMale    Gender = 1
	GenderFemale  Gender = 2
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;uniqueIndex;not null;type:varchar(18)"`
	Password string `json:"password" gorm:"column:password;not null;type:varchar(18)"`
	Email    string `json:"email" gorm:"column:email;uniqueIndex;not null;type:varchar(40)"`
	Gender   Gender `json:"gender" gorm:"column:gender;type:tinyint;default:0"`
	Age      uint8  `json:"age" gorm:"column:age;type:tinyint;not null"`
}

func (User) TableName() string {
	return "bms_user"
}
