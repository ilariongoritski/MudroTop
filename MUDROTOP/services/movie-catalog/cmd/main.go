package main

import (
	"context"
	"log"

	"mudrotop/services/movie-catalog/app"
)

func main() {
	if err := app.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
