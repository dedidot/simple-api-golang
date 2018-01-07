package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"./controllers"
	"../config"
	"../migrate"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = migrate.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Post("/book", a.InputBook)
	a.Get("/book", a.ListBook)
	a.Get("/book/{codes:[1-9]+}", a.OneBook)
	a.Put("/book/{codes:[1-9]+}", a.UpdateBook)
	a.Delete("/book/{codes:[1-9]+}", a.DeletedBook)
}

//handler method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Get")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Put")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("Delete")
}

//get data
func (a *App) InputBook(w http.ResponseWriter, r *http.Request) {
	controllers.InputBook(a.DB, w, r)
}

func (a *App) ListBook(w http.ResponseWriter, r *http.Request) {
	controllers.ListBook(a.DB, w, r)
}

func (a *App) OneBook(w http.ResponseWriter, r *http.Request) {
	controllers.OneBook(a.DB, w, r)
}

func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateBook(a.DB, w, r)
}
 
func (a *App) DeletedBook(w http.ResponseWriter, r *http.Request) {
	controllers.DeletedBook(a.DB, w, r)
} 

//run
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}