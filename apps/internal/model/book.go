package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title        string    `json:"title" gorm:"column:title;uniqueIdex:uidx_bms_book_title;not null;type:varchar(20)"`
	Author       uint64    `json:"author" gorm:"column:author_id;not null;type:bigint"`
	Summary      string    `json:"summary" gorm:"column:summary;"`
	Price        int32     `json:"price" gorm:"column:price;not null;type:int"`
	Publisher    uint64    `json:"publisher" gorm:"column:publisher;type:bigint;not null"`
	Publish_time time.Time `json:"publish_time" gorm:"column:publish_time;not null;type:datetime"`
}
