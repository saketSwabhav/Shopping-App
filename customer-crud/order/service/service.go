package service

import (
	"customerCrud/model/order"
	"customerCrud/repository"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	Repository repository.Repository
	Db         *gorm.DB
}

func NewService(re *repository.Repository, d *gorm.DB) *Service {
	return &Service{
		Repository: *re,
		Db:         d,
	}
}
func (s *Service) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")

}

func (s *Service) GetAllOrders(orders *[]order.Order, customerID uuid.UUID) error {

	uow := repository.NewUnitOfWork(s.Db, true)

	// err := db.Where("email = ?", email).First(&result)
	fmt.Println(orders)
	var queryProcessors []repository.QueryProcessor
	queryProcessors = append(queryProcessors, repository.Where("customer_id=?", customerID))
	queryProcessors = append(queryProcessors, repository.Preload([]string{"Product"}))
	err := s.Repository.Find(uow, &orders, queryProcessors)
	if err != nil {
		return err
	}
	// err := db.Where("email = ?", email).First(&result)

	return nil
}
func (s *Service) GetAll(customers *[]order.Order) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.GetAll(uow, &customers, "Orders")
}

func (s *Service) CreateNewOrder(orders *order.Order) error {

	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Add(uow, &orders)
	if err != nil {
		uow.DB.Rollback()
		return err

	}
	uow.Commit()
	return nil
}
func (s *Service) UpdateOrder(Order *order.Order) error {

	// fmt.Println(Order.ID)
	uow := repository.NewUnitOfWork(s.Db, false)
	err := s.Repository.Save(uow, &Order)
	if err != nil {
		uow.DB.Rollback()
		return err
	}
	uow.Commit()
	return nil
}

func (s *Service) ReturnSingleOrder(Order *order.Order, id uuid.UUID) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.Get(uow, &Order, id, []string{"Product"})
}

func (s *Service) DeleteOrder(Order *order.Order) {

	uow := repository.NewUnitOfWork(s.Db, true)
	s.Repository.Delete(uow, &Order)
}
