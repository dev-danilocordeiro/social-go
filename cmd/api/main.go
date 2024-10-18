package main

import (
	"log"

	"github.com/danilocordeirodev/social-go/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()

	log.Fatal(app.run(mux))

}
