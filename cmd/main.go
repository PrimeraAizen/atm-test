package main

import (
	"ATMtesttask"
	"log"
)

func main() {
	srv := new(ATMtesttask.Server)
	err := srv.Run("8080", nil)
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
