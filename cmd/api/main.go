package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()

	log.Fatal(app.run(mux))

}
