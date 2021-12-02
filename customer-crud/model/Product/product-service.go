package product

import (
	"customerCrud/repository"

	"github.com/jinzhu/gorm"
)

type ProductService struct {
	Repository repository.Repository
	Db         *gorm.DB
}

func NewService(re *repository.Repository, d *gorm.DB) *ProductService {
	return &ProductService{
		Repository: *re,
		Db:         d,
	}
}

func (s *ProductService) GetAll(products *[]Product) error {

	uow := repository.NewUnitOfWork(s.Db, true)
	err := s.Repository.GetAll(uow, &products)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) CreateCustomer(customer *Product) error {

	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Add(uow, &customer)
	if err != nil {
		uow.DB.Rollback()
		return err

	}
	uow.Commit()
	return nil
}
