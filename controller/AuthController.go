package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func RegisterCompany(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/register_company.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	err = t.ExecuteTemplate(w, "register_company", nil)
}
