package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Tittle    string `json:"title" form:"title"`
	Publisher string `json:"publisher" form:"publisher"`
}
