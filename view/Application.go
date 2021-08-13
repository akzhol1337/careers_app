package view

import (
	"careers_tool/controller"
	_ "database/sql"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "net/http"
)

func HandleFunc(){
	route := mux.NewRouter()
	route.HandleFunc("/", controller.Index).Methods("GET")
}
