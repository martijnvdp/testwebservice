package main

import (
	"fmt"

	"github.com/martijnxd/testwebservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Martijn",
		LastName:  "Last name",
	}
	fmt.Println(u)
}
