package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	user := models.User{
		ID:        2,
		FirstName: "Spwn",
		LastName:  "MC",
	}
	fmt.Println(user)
}
