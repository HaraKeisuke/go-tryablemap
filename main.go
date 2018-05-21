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

	fmt.Println(result)

	data2 := map[string][]interface{}{
		"sample": []interface{}{
			"first", "second", "third",
		},
	}
	_map2 := tryablemap.NewTryableArray(data2)
	result2 := _map2.Try("sample").Value(0)

	fmt.Println(result2)
}
