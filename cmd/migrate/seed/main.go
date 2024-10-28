package main

import (
	"log"

	"github.com/danilocordeirodev/social-go/internal/db"
	"github.com/danilocordeirodev/social-go/internal/store"
)

func main() {
	conn, err := db.New("postgres://user:adminpassword@localhost/social?sslmode=disable", 3, 3, "15m")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
