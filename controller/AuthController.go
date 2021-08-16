package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"careers_tool/model/repositories"
)

type userController struct {
	userRepository repositories.userRepository
}

func RegisterCompany(w http.ResponseWriter, r *http.Request){
	userRepository = repositories.UserRepository()
	t, err := template.ParseFiles("templates/register_company.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	err = t.ExecuteTemplate(w, "register_company", nil)
}

