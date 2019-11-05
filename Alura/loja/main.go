package main

import (
	"net/http"

	"github.com/fabiossilva00/Go/Alura/loja/routers"
)

func main() {
	routers.LoadingRouters()
	http.ListenAndServe(":8000", nil)
}
