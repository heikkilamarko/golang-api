package main

import (
	"os"
	"products-api/app"
	"products-api/app/config"

	"github.com/rs/zerolog"
)

func main() {
	c := config.New()
	c.Load()

	l := zerolog.New(os.Stderr).With().Timestamp().Logger()

	a := app.New(c, &l)
	a.Run()
}
