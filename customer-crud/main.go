package main

import (
	"context"
	"customerCrud/customer/controller"
	"customerCrud/customer/service"
	product "customerCrud/model/Product"
	"customerCrud/model/customer"
	"customerCrud/model/order"
	"customerCrud/model/user"
	orderController "customerCrud/order/controller"
	orderService "customerCrud/order/service"
	"customerCrud/repository"
	"customerCrud/security/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to Open", err)
		return
	}
	log.Println("Connection Established")

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&customer.Customer{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&product.Product{})
	db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	db.Model(&order.Order{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	myRouter := mux.NewRouter().StrictSlash(true)
	repo := repository.NewRepository()

	// myRouter.Use(middleware.Middleware)

	// fmt.Println(security.GenerateToken())
	// log.Fatal(http.ListenAndServe(":10000", myRouter))

	middlewareRouter := myRouter.PathPrefix("/").Subrouter()
	getRouter := myRouter.PathPrefix("/").Subrouter()
	middlewareRouter.Use(middleware.Middleware)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Token"})
	// headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Total-Count", "token", "totalLifetimeValue"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origin := handlers.AllowedOrigins([]string{"*"})

	server := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(myRouter),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}

	RegisterUserRoutes(getRouter, middlewareRouter, repo, db)
	var wait time.Duration

	go func() {
		log.Fatal(server.ListenAndServe())
		fmt.Println(" ======= Listening at port :8080")
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	func() {
		fmt.Println("Closing DB")
		db.Close()
	}()
	fmt.Println("Server ShutDown.......")
	os.Exit(0)

}

func RegisterUserRoutes(getRouter, middlewareRouter *mux.Router, repo repository.Repository, db *gorm.DB) {
	services := service.NewService(&repo, db)
	control := controller.NewController(services)
	control.HandleRequests(getRouter, middlewareRouter)
	productService := product.NewService(&repo, db)
	productController := *product.NewController(productService)
	productController.HandleRequests(getRouter, middlewareRouter)
	// user.NewController()
	userService := user.NewService(&repo, db)
	userController := user.NewController(userService)
	userController.HandleRequests(getRouter, middlewareRouter)

	orderSer := orderService.NewService(&repo, db)
	orderCont := orderController.NewController(orderSer)
	orderCont.HandleRequests(getRouter)

	// panic("unimplemented")
}
