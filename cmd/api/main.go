package main

import (
	"log"

	"github.com/danilocordeirodev/social-go/internal/env"
	"github.com/danilocordeirodev/social-go/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()

	log.Fatal(app.run(mux))

}
