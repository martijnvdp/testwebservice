package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/martijnxd/testwebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	path := r.URL.Path[1:]
	if strings.HasSuffix(path, "@testing.com") {
		fmt.Fprintf(w, "Hello tester %s\n", strings.TrimSuffix(path, "@gtesting.com"))
		return
	}
	fmt.Fprintf(w, "Hello %s\n<br><br>", path)
	fmt.Fprintf(w, "<a href=\"users\">List users</a>")
	fmt.Fprintf(w, "<br><br>Test web app with GO GO")
}
