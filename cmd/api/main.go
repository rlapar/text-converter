package main

import (
	"log"
	"os"
	apipkg "text-converter/internal/api"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	api := apipkg.NewAPI()
	err := api.ListenAndServe(":8080")
	if err != nil {
		log.Printf("Server failed", err.Error())
		exitCode = 1
		return
	}
}