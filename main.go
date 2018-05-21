package main

import (
	"fmt"

	tryablemap "github.com/HaraKeisuke/go-tryablemap/lib"
)

func main() {
	data := map[string]interface{}{
		"first": map[string]int{
			"second": 123,
		},
	}
	_map := tryablemap.NewTryableMap(data)
	result := _map.Try("first").Value("second")

	fmt.Println(result)

	data2 := map[string][]string{
		"sample": []string{
			"first", "second", "third",
		},
	}

	_map2 := tryablemap.NewTryableMap(data2)
	result2 := _map2.TryArray("sample").Value(0)

	fmt.Println(result2)

	data3 := [][]int{
		[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9},
	}
	_map3 := tryablemap.NewTryableArray(data3)
	result4 := _map3.TryArray(0).Value(2)

	fmt.Println(result4)
}
