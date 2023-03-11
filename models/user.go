package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 					string		`json:"name"`
	Surname 				string		`json:"surname"`
	Email					string 		`json:"email"`
	Password				string		`json:"password"`
	Available_funds 	float64		`json:"available_funds"`
	TotalFunds			float64		`json:"total_funds"`
	Lang 					string		`json:"lang"`
	Theme 				string		`json:"theme"`
}

