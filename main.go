package main

import (
	"net/http"

	"github.com/martijnxd/testwebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
