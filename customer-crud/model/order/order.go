package order

import (
	"customerCrud/model"
	product "customerCrud/model/Product"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	model.Model
	ProductID  uuid.UUID       `gorm:"ForeignKey:ProductID;type:varchar(36)" json:"productId"`
	CustomerID uuid.UUID       `gorm:"ForeignKey:CustomerID;type:varchar(36)" json:"customerId"`
	Quantity   int             `gorm:" type:int" json:"quantity"`
	IsPaid     *bool           `gorm:" type:tinyint" json:"isPaid"`
	OrderDate  time.Time       `json:"orderDate"`
	Product    product.Product `json:"products"`
}
