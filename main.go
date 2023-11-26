package main

import (
	"log"
	"net/http"

	"github.com/06202003/apiInventory/middlewares"

	"github.com/06202003/apiInventory/controllers/authcontroller"
	"github.com/06202003/apiInventory/controllers/categorycontroller"
	"github.com/06202003/apiInventory/controllers/employeecontroller"
	"github.com/06202003/apiInventory/controllers/inventorycontroller"
	"github.com/06202003/apiInventory/controllers/reporthistorypemakaiancontroller"
	"github.com/06202003/apiInventory/controllers/reporthistoryperbaikancontroller"
	"github.com/06202003/apiInventory/controllers/roomcontroller"
	"github.com/06202003/apiInventory/controllers/locationcontroller"
	"github.com/06202003/apiInventory/controllers/usagecontroller"
	"github.com/06202003/apiInventory/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	
	api.HandleFunc("/categories", categorycontroller.Index).Methods("GET")
	api.HandleFunc("/categories/{id_kategori}", categorycontroller.Show).Methods("GET")
	api.HandleFunc("/categories", categorycontroller.Create).Methods("POST")
	api.HandleFunc("/categories/{id_kategori}", categorycontroller.Update).Methods("PUT")
	api.HandleFunc("/categories/{id_kategori}", categorycontroller.Delete).Methods("DELETE")

	api.HandleFunc("/employees", employeecontroller.Index).Methods("GET")
	api.HandleFunc("/employees/{nomor_induk}", employeecontroller.Show).Methods("GET")
	api.HandleFunc("/employees", employeecontroller.Create).Methods("POST")
	api.HandleFunc("/employees/{nomor_induk}", employeecontroller.Update).Methods("PUT")
	api.HandleFunc("/employees/{nomor_induk}", employeecontroller.Delete).Methods("DELETE")

	api.HandleFunc("/inventories", inventorycontroller.Index).Methods("GET")
	api.HandleFunc("/inventories/{kode_aset}", inventorycontroller.Show).Methods("GET")
	api.HandleFunc("/inventories", inventorycontroller.Create).Methods("POST")
	api.HandleFunc("/inventories/{kode_aset}", inventorycontroller.Update).Methods("PUT")
	api.HandleFunc("/inventories/{kode_aset}", inventorycontroller.Delete).Methods("DELETE")

	api.HandleFunc("/usages", usagecontroller.Index).Methods("GET")
	api.HandleFunc("/usages/{id_pemakaian}", usagecontroller.Show).Methods("GET")
	api.HandleFunc("/usages", usagecontroller.Create).Methods("POST")
	api.HandleFunc("/usages/{id_pemakaian}", usagecontroller.Update).Methods("PUT")
	api.HandleFunc("/usages/{id_pemakaian}", usagecontroller.Delete).Methods("DELETE")


	api.HandleFunc("/usageHistories", reporthistorypemakaiancontroller.Index).Methods("GET")
	api.HandleFunc("/usageHistories/{id}", reporthistorypemakaiancontroller.Show).Methods("GET")

	api.HandleFunc("/repairHistories", reporthistoryperbaikancontroller.Index).Methods("GET")
	api.HandleFunc("/repairHistories/{id}", reporthistoryperbaikancontroller.Show).Methods("GET")
	api.HandleFunc("/repairHistories", reporthistoryperbaikancontroller.Create).Methods("POST")
	api.HandleFunc("/repairHistories/{id}", reporthistoryperbaikancontroller.Update).Methods("PUT")
	api.HandleFunc("/repairHistories/{id}", reporthistoryperbaikancontroller.Delete).Methods("DELETE")

	api.HandleFunc("/rooms", roomcontroller.Index).Methods("GET")
	api.HandleFunc("/rooms/{id_ruangan}", roomcontroller.Show).Methods("GET")
	api.HandleFunc("/rooms", roomcontroller.Create).Methods("POST")
	api.HandleFunc("/rooms/{id_ruangan}", roomcontroller.Update).Methods("PUT")
	api.HandleFunc("/rooms/{id_ruangan}", roomcontroller.Delete).Methods("DELETE")

	api.HandleFunc("/locations", locationcontroller.Index).Methods("GET")
	api.HandleFunc("/locations/{id_lokasi}", locationcontroller.Show).Methods("GET")
	api.HandleFunc("/locations", locationcontroller.Create).Methods("POST")
	api.HandleFunc("/locations/{id_lokasi}", locationcontroller.Update).Methods("PUT")
	api.HandleFunc("/locations/{id_lokasi}", locationcontroller.Delete).Methods("DELETE")

	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
