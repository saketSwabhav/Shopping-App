package customer

import (
	"customerCrud/model"
	"customerCrud/model/order"
	"time"
)

type Customer struct {
	model.Model
	Email    string        `gorm:" type:varchar(100)" json:"email"`
	UserPass string        `gorm:" type:varchar(100)" json:"userPass"`
	Fname    string        `gorm:" type:varchar(20)" json:"fName"`
	Lname    string        `gorm:" type:varchar(20)" json:"lName"`
	Age      int           `gorm:" type:int" json:"age"`
	IsMale   *bool         `gorm:" type:tinyint" json:"isMale"`
	DOB      time.Time     `json:"dob"`
	Orders   []order.Order `json:"orders"`
}
