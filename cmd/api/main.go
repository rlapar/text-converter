package main

import (
	"os"
	apipkg "text-converter/internal/api"
	"text-converter/internal/cfg"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	api := apipkg.NewAPI()
	err := api.ListenAndServe(":8080")
	if err != nil {
		cfg.Logger.Info("Server failed", err.Error())
		exitCode = 1
		return
	}
}