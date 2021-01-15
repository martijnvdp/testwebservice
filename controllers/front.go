package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers function
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
	http.HandleFunc("/adduser", AddUserGui)
	http.HandleFunc("/addusers", AddUserGui)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
