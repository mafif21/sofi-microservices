package main

import (
	"github.com/joho/godotenv"
	"pendaftaran-sidang/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app.StartApp()
}
