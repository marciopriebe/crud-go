package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string // seria level junior, pleno, 
	Company string
	Location string
	Remote bool
	Link string
	Salary int64
}