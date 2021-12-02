package product

import "customerCrud/model"

type Product struct {
	model.Model
	ItemName  string `gorm:" type:varchar(100)" json:"itemName"`
	ItemDesc  string `gorm:" type:varchar(100)" json:"itemDesc"`
	ItemPrice int    `gorm:" type:int" json:"itemPrice"`
}
