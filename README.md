# Go TryableMap

Tryable Map is safety dig tool for Map(key/value).

For example, you can chain like Ruby's try method and find the property.

### How to use
```
package main

import (
  "fmt"
  tryablemap "github.com/HaraKeisuke/go-tryablemap/lib"
)

func main() {
  sample := map[string]interface{}{
    "first": map[string]interface{}{
      "second": 123,
    },
  }
  _map := tryablemap.NewTryableMap(sample)
  result := _map.Try("first").Value("second")

  fmt.Print(result) //-> 123
}
```

### Finally
Please tell me if there is a better way :)