package main

import (
	"products-api/app"
)

func main() {
	a := app.NewApp()
	a.Initialize()
	a.Run()
}
