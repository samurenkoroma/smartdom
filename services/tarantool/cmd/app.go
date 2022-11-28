package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
)

func main() {
	opts := tarantool.Opts{User: "guest"}
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
	}
	resp, err := conn.Insert(999, []interface{}{99999, "BB"})
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("Code", resp.Code)
	}
}
