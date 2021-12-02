package controller

import (
	"customerCrud/model/order"
	"customerCrud/order/service"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Controller struct {
	orderService service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		orderService: *s,
	}
}
func (h *Controller) HandleRequests(r *mux.Router) {
	// creates a new instance of a mux router
	// myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	r.HandleFunc("/", h.orderService.HomePage)
	r.HandleFunc("/orders", h.GetAll)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	r.HandleFunc("/order", h.CreateNeworder).Methods("POST")
	r.HandleFunc("/order/{id}", h.ReturnSingleorderomer).Methods("GET")
	r.HandleFunc("/order/{id}", h.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/order/{id}", h.Updateorder).Methods("PUT")
	r.HandleFunc("/{customerID}/orders", h.GetAllOrders).Methods("GET")

}

func (h *Controller) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["customerID"]
	order := []order.Order{}
	custID, err := uuid.FromString(key)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	err = h.orderService.GetAllOrders(&order, custID)

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	userOrder, err := json.Marshal(&order)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	fmt.Println(order)
	w.Write(userOrder)

}
func (h *Controller) GetAll(w http.ResponseWriter, r *http.Request) {

	order := []order.Order{}
	h.orderService.GetAll(&order)
	data, err := json.Marshal(&order)
	if err != nil {
		fmt.Fprint(w, errors.New("internal error"))
		return
	}

	fmt.Fprint(w, string(data))
}

func (h *Controller) CreateNeworder(w http.ResponseWriter, r *http.Request) {
	order := order.Order{}
	// Unmarshal json.
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &order)
	// (r, &order)
	if err != nil {
		// log.NewLogger().Error(err.Error())
		fmt.Fprint(w, "Error in adding ", err)
		// web.RespondError(w, errors.NewHTTPError("unable to parse requested data", http.StatusBadRequest))
		return
	}
	h.orderService.CreateNewOrder(&order)
	fmt.Fprint(w, "Record Added Successfully")

}
func (h *Controller) Updateorder(w http.ResponseWriter, r *http.Request) {
	order := order.Order{}
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
	err := json.Unmarshal(body, &order)
	// (r, &order)
	if err != nil {
		// log.NewLogger().Error(err.Error())
		// web.RespondError(w, errors.NewHTTPError("unable to parse requested data", http.StatusBadRequest))
		return
	}
	order.ID = id
	errs := h.orderService.UpdateOrder(&order)
	if errs != nil {
		fmt.Fprint(w, "Errors updating values ", errs)
	}

}
func (h *Controller) ReturnSingleorderomer(w http.ResponseWriter, r *http.Request) {
	order := order.Order{}

	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
	}
	id, err := uuid.FromString(input)

	if err != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}

	h.orderService.ReturnSingleOrder(&order, id)
	data, err := json.Marshal(&order)
	if err != nil {
		fmt.Fprint(w, errors.New("internal error"))
		return
	}

	fmt.Fprint(w, string(data))

}
func (h *Controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {

	order := order.Order{}
	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
	}
	id, err := uuid.FromString(input)

	if err != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}
	order.ID = id
	h.orderService.DeleteOrder(&order)

}
