package main

import (
	"context"
	"log"
	"smartdom/pkg/store/mongodb"
	"smartdom/services/library/internal/delivery/http"
	"smartdom/services/library/internal/repository/storage/mongo"
	"smartdom/services/library/internal/useCase/book"
)

func main() {
	client, err := mongodb.NewClient(context.Background(), mongodb.Settings{
		Host:     "localhost",
		Port:     27017,
		Database: "smartdom",
		User:     "smartdom",
		Password: "smartdom",
	})
	if err != nil {
		log.Fatal(err)
	}
	repository := mongo.New(client)
	useCase := book.New(repository)
	delivery := http.New(useCase)
	if err := delivery.Run(); err != nil {
		return
	}
	//

}
