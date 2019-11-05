package routers

import (
	"net/http"

	"github.com/fabiossilva00/Go/Alura/loja/controllers"
)

func LoadingRouters() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.InsertProduts)
	http.HandleFunc("/delete", controllers.DeleteProdut)
	http.HandleFunc("/edit", controllers.EditProdut)
	http.HandleFunc("/update", controllers.UpdateProdut)
}
