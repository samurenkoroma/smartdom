package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Order struct {
	ID       int
	Products string
}

func (o Order) String() string {
	return fmt.Sprintf("id:%d;products:%s", o.ID, o.Products)
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

	client.Set("hashawd", "Order{ID:0,Products:'awdadw'}", time.Hour*200)
	fmt.Printf("from redis: %s\n", client.GeoDist("city", "Odessa", "Piter1", "km"))
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
