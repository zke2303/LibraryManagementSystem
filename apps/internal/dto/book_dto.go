package dto

import "time"

type CreateBookRequest struct {
	Title        string    `json:"title" binding:"required,min=2,max=20"`
	Author       uint64    `json:"author" binding:"required"`
	Summary      string    `json:"summary"`
	Price        uint32    `json:"price" binding:"required,min=1,max=9999999"`
	Publisher    uint64    `json:"publisher" binding:"required"`
	Publish_time time.Time `json:"publish_time" binding:"required"`
	ISBN         string    `json:"ISBN" binding:"required,min=13,max=13"`
}
