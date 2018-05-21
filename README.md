# Go TryableMap

Tryable Map is safety dig tool for Map(key/value) and Array.

For example, you can chain like Ruby's try method and find the property.

### How to use

```
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
```

### Todo
- [ ] Convert TryableArray to TryableMap
- [ ] Convert TryableMap to TryableArray

### Finally

Please tell me if there is a better way :)
