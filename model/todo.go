package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required,oneof=done pending" `
}
