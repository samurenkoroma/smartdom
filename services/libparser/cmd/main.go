package main

import (
	"context"
	"fmt"
	"log"
	"smartdom/pkg/store/mongodb"
	"smartdom/services/libparser/internal/delivery/cli"
	"smartdom/services/libparser/internal/repository/storage/mongodbrepo"
	"smartdom/services/libparser/internal/useCase/parser"
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
	cli.New(parser.New(mongodbrepo.New(client)), cli.Options{
		Dir: "/home/itzov/books",
	}).Parse()
	fmt.Println(client)
}
