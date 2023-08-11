package main

import (
	"log"

	"github.com/ByteNinja42/ExpensesTool/internal/entities"
	"github.com/ByteNinja42/ExpensesTool/internal/service"
)

func main() {
	srv := service.NewExpensesService()
	err := srv.UserSignUp(entities.UserSignUp{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Password:  "",
	})
	if err != nil {
		log.Fatal(err)
	}
}
