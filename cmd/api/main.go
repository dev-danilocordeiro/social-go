package main

import (
	"log"

	"github.com/danilocordeirodev/social-go/internal/db"
	"github.com/danilocordeirodev/social-go/internal/env"
	"github.com/danilocordeirodev/social-go/internal/store"
)

const version = "0.0.1"

//	@title			Sociale API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@secyrityDefinitions.apiKey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "locahost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("server has started at %s", app.config.addr)

	mux := app.mount()

	log.Fatal(app.run(mux))

}
