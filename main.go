package main

import (
	"fmt"

	tryablemap "github.com/HaraKeisuke/go-tryablemap/lib"
)

func main() {
	data := map[string]interface{}{
		"first": map[string]interface{}{
			"second": 123,
		},
	}
	_map := tryablemap.NewTryableMap(data)
	result := _map.Try("first").Value("second")

	fmt.Print(result)
}
