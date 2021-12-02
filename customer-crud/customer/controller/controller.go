package controller

import (
	"customerCrud/customer/service"
	"customerCrud/model/customer"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Controller struct {
	CustService service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		CustService: *s,
	}
}
func (h *Controller) HandleRequests(router, middleware *mux.Router) {
	router.HandleFunc("/", h.CustService.HomePage)
	middleware.HandleFunc("/customers", h.GetAll)
	middleware.HandleFunc("/customer", h.CreateCustomer).Methods("POST")
	middleware.HandleFunc("/customer/{id}", h.GetCustomer).Methods("GET")
	middleware.HandleFunc("/customer/{id}", h.DeleteCustomer).Methods("DELETE")
	middleware.HandleFunc("/customer/{id}", h.UpdateCustomer).Methods("PUT")

}

func (h *Controller) GetAll(w http.ResponseWriter, r *http.Request) {

	cust := []customer.Customer{}
	err := h.CustService.GetAll(&cust)
	if err != nil {
		fmt.Fprint(w, errors.New(err.Error()))
		return
	}
	data, err := json.Marshal(&cust)
	if err != nil {
		fmt.Fprint(w, errors.New(err.Error()))
		return
	}

	w.Write(data)
	// fmt.Fprint(w, string(data))
}

func (h *Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	cust := customer.Customer{}
	// Unmarshal json.
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &cust)
	fmt.Fprint(w, &cust)

	if err != nil {
		fmt.Fprint(w, "Error in adding ", err)
		return
	}
	err = h.CustService.CreateCustomer(&cust)
	if err != nil {
		fmt.Fprint(w, "Error in adding ", err)
		return
	}
	w.Write([]byte("record added successfully"))
	// fmt.Fprint(w, "Record Added Successfully")

}
func (h *Controller) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	cust := customer.Customer{}
	// Unmarshal json.
	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
	}
	id, er := uuid.FromString(input)

	if er != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &cust)
	if err != nil {
		return
	}
	cust.ID = id
	errs := h.CustService.UpdateCustomer(&cust)
	if errs != nil {
		fmt.Fprint(w, "Errors updating values ", errs)
		return
	}
	w.Write([]byte("record updated successfully"))

	fmt.Fprint(w, "Record Updated Successfully")

}
func (h *Controller) GetCustomer(w http.ResponseWriter, r *http.Request) {
	cust := customer.Customer{}

	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
		return
	}
	id, err := uuid.FromString(input)

	if err != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}

	errs := h.CustService.GetCustomer(&cust, id)
	data, err := json.Marshal(&cust)
	if err != nil {
		fmt.Fprint(w, errors.New("internal error"))
		return
	}
	if errs != nil {
		fmt.Fprint(w, "error: ", errs)
		return
	}

	w.Write(data)

	fmt.Fprint(w, string(data))

}
func (h *Controller) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	cust := customer.Customer{}
	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
	}
	id, err := uuid.FromString(input)

	if err != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}
	cust.ID = id
	errs := h.CustService.DeleteCustomer(&cust)
	if errs != nil {
		fmt.Fprint(w, "error: ", errs)
		return
	}
	w.Write([]byte("record deleted successfully"))

	fmt.Fprint(w, "Deleted successfully ")

}
