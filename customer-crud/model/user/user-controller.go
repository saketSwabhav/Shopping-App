package user

import (
	"customerCrud/model/customer"
	"customerCrud/security/auth"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Service *UserService
}

func NewController(service *UserService) *Controller {
	return &Controller{
		Service: service,
	}
}
func (h *Controller) HandleRequests(router, middleware *mux.Router) {

	router.HandleFunc("/login", h.UserLogin).Methods("POST")
	router.HandleFunc("/register", h.Register).Methods("POST")

}

//Register will register the user.
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" =========================== REGISTER =========================== ")
	var user = &customer.Customer{}
	var err error

	userDetails, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Response could not be read"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(userDetails, user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPass), 8)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.UserPass = string(hashPassword)

	err = c.Service.Add(user)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("Error while adding user, " + err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("User ID -> ", user.ID)
	tokenDetails, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("after login", err)
		return
	}
	fmt.Println(tokenDetails)
	w.Header().Set("Content-Type", "application/json")

	newUser, err := json.Marshal(User{
		Token: tokenDetails,
		ID:    user.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(string(newUser))
	w.Write(newUser)
}

func (c *Controller) UserLogin(w http.ResponseWriter, r *http.Request) {

	var loginUser = &customer.Customer{}
	// var validateUser = &customer.Customer{}
	var err error

	log.Println(" ---------------------- Inside userlogin ---------------------- ")

	userDetails, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(userDetails, loginUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userPass := loginUser.UserPass
	err = c.Service.Get(loginUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(loginUser.UserPass), []byte(userPass)); err != nil {
		http.Error(w, "Username or password is invalid", http.StatusUnauthorized)
		log.Println("Username or password is invalid password", err)
		return
	}
	// fmt.Println(" loginusername -> ", loginUser.Username)
	// fmt.Println(" validateusername -> ", validateUser.Username)
	fmt.Println(loginUser.ID)
	tokenDetails, err := auth.GenerateToken(loginUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("after login", err)
		return
	}
	fmt.Println(tokenDetails)
	w.Header().Set("Content-Type", "application/json")

	user, err := json.Marshal(User{
		Token: tokenDetails,
		ID:    loginUser.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(string(user))
	w.Write(user)
}
