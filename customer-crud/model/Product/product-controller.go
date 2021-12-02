package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	ProductService ProductService
}

func NewController(s *ProductService) *Controller {
	return &Controller{
		ProductService: *s,
	}
}

func (h *Controller) HandleRequests(router, middleware *mux.Router) {
	// router.HandleFunc("/", h.ProductService.HomePage)
	router.HandleFunc("/products", h.GetAll)
	router.HandleFunc("/product", h.CreateProduct).Methods("POST")
	// middleware.HandleFunc("/customer/{id}", h.GetCustomer).Methods("GET")
	// middleware.HandleFunc("/customer/{id}", h.DeleteCustomer).Methods("DELETE")
	// middleware.HandleFunc("/customer/{id}", h.UpdateCustomer).Methods("PUT")

}
func (h *Controller) GetAll(w http.ResponseWriter, r *http.Request) {

	product := []Product{}
	err := h.ProductService.GetAll(&product)
	if err != nil {
		fmt.Fprint(w, errors.New(err.Error()))
		return
	}
	data, err := json.Marshal(&product)
	if err != nil {
		fmt.Fprint(w, errors.New(err.Error()))
		return
	}

	w.Write(data)
	// fmt.Fprint(w, string(data))
}

func (h *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{}
	// Unmarshal json.
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &product)
	fmt.Fprint(w, &product)

	if err != nil {
		fmt.Fprint(w, "Error in adding ", err)
		return
	}
	err = h.ProductService.CreateCustomer(&product)
	if err != nil {
		fmt.Fprint(w, "Error in adding ", err)
		return
	}
	w.Write([]byte("record added successfully"))
	// fmt.Fprint(w, "Record Added Successfully")

}
