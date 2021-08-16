package view

import (
	"careers_tool/controller"
	_ "database/sql"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http"
)

func HandleFunc(){
	route := mux.NewRouter()
	route.HandleFunc("/", controller.Index).Methods("GET")
	route.HandleFunc("/register", controller.RegisterCompany).Methods("GET")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", route)
	http.ListenAndServe(":8080", nil)
}
