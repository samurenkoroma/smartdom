package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Order struct {
	ID       int       `json:"id,omitempty"`
	Products string    `json:"products,omitempty"`
	Created  time.Time `json:"created"`
}

func main() {
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	client.GeoAdd("city", &redis.GeoLocation{
		Name:      "Piter1",
		Longitude: 59,
		Latitude:  30.315790,
	}, &redis.GeoLocation{
		Name:      "Odessa",
		Longitude: 60,
		Latitude:  30.315790,
	})

	client.Set("hash", Order{
		ID:       0,
		Products: "awdadw",
		Created:  time.Time{},
	}, time.Hour*200)
	fmt.Printf("from redis: %s\n", client.GeoDist("city", "Odessa", "Piter1", "km"))
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
