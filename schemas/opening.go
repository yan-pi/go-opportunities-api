package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	//Discription string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   string
}
