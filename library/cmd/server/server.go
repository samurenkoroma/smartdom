package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"smarthome/library/internal/config"
	"smarthome/library/internal/domain/book"
	"smarthome/library/internal/domain/book/db"
	"smarthome/library/pkg/client/mongodb"
)

func main() {

	cfg := config.GetConfig()
	ctx := context.Background()
	client, err := mongodb.NewClient(ctx, cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Database)
	if err != nil {
		log.Fatal(err)
	}
	service := book.NewService(db.NewStorage(client, "Books"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		books, err := service.GetAll(ctx)
		if err != nil {
			return
		}
		if err = json.NewEncoder(writer).Encode(books); err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
