package main

import (
	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.routesfunc(r)

}
