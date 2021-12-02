package user

import (
	"customerCrud/model/customer"
	"customerCrud/repository"
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	Repository repository.Repository
	Db         *gorm.DB
}

func NewService(re *repository.Repository, d *gorm.DB) *UserService {
	return &UserService{
		Repository: *re,
		Db:         d,
	}
}

func (s *UserService) Get(customer *customer.Customer) error {

	uow := repository.NewUnitOfWork(s.Db, true)

	// err := db.Where("email = ?", email).First(&result)
	fmt.Println(customer)
	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Where("email=?", customer.Email))
	err := s.Repository.Find(uow, &customer, queryProcessors)
	if err != nil {
		return err
	}
	// err := db.Where("email = ?", email).First(&result)

	return nil
}
func (s *UserService) Add(customer *customer.Customer) error {

	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Add(uow, &customer)
	if err != nil {
		uow.DB.Rollback()
		return err

	}
	uow.Commit()
	return nil
}
