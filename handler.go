package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Watch struct {
	gorm.Model
	Name string
	Price string
}

func Conn() *gorm.DB {
	// cenect to database sqlite
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Watch{})

	return db
}

func MuxEngine() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/create", Create)
	mux.HandleFunc("/read", Read).Methods("GET")
	mux.HandleFunc("/update", Update).Methods("PUT")
	mux.HandleFunc("/delete/{id}", Delete).Methods("DELETE")
	return mux
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	// get form data
	name := r.FormValue("name")
	price := r.FormValue("price")

	// connect to database
	db := Conn()

	// create data
	db.Create(&Watch{Name: name, Price: price})

	// status code
	w.WriteHeader(http.StatusOK)
	// type json
	w.Header().Set("Content-Type", "application/json")
	// response
	w.Write([]byte(`{"name": "` + name + `", "price": "` + price + `"}`))
}

func Read(w http.ResponseWriter, r *http.Request) {
	// connect to database
	db := Conn()

	var watches []Watch
	// read data
	db.Find(&watches)

	// status code
	w.WriteHeader(http.StatusOK)
	// type json
	w.Header().Set("Content-Type", "application/json")
	// response success
	w.Write([]byte(`{"status": "success"}`))
	
}

func Update(w http.ResponseWriter, r *http.Request) {
	// get form data
	name := r.FormValue("name")
	price := r.FormValue("price")

	// connect to database
	db := Conn()

	// update data
	db.Model(&Watch{}).Where("name = ?", name).Update("price", price)

	// status code
	w.WriteHeader(http.StatusOK)
	// type json
	w.Header().Set("Content-Type", "application/json")
	// response
	w.Write([]byte(`{"name": "` + name + `", "price": "` + price + `"}`))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// get form url
	vars := mux.Vars(r)
	id := vars["id"]


	// connect to database
	db := Conn()

	var watch Watch
	// delete data
	db.Delete(&watch, id)

	// status code
	w.WriteHeader(http.StatusOK)
	// type json
	w.Header().Set("Content-Type", "application/json")
	// response
	w.Write([]byte(`{"id": "` + id + `"}`))
}