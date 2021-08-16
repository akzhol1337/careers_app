package main

import (
	"careers_tool/view"
	"careers_tool/model"
	"careers_tool/model/repositories"
)

func main(){
	db := model.GetConnection("srv-pleskdb26.ps.kz", "3306", "oqschool_golang", "oqschool_admin", "A0y!zp33")
	userRepositoy := repositories.UserRepository(db)

	view.HandleFunc()
}

