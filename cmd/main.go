package main

import "github.com/DexScen/api.calculator/internal/application"

func main() {
	app := application.New()
	app.RunServer()
}