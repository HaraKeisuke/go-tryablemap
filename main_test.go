package main

import (
	"testing"

	"github.com/HaraKeisuke/go-tryablemap/lib"
)

func TestTryableMap(t *testing.T) {
	data := map[string]interface{}{
		"first": map[string]interface{}{
			"second": 123,
			"second_2": map[string]interface{}{
				"third": "hogehoge",
			},
		},
	}
	tryable := tryablemap.NewTryableMap(data)
	second := tryable.Try("first").Value("second")
	if second != 123 {
		t.Errorf("Second is not invalid")
	}

	third := tryable.Try("first").Try("second_2").Value("third")
	if third != "hogehoge" {
		t.Errorf("Second is not invalid")
	}

	invalidKey := tryable.Try("first").Try("second_0").Value("NOT_EXIST_KEY")
	if invalidKey != nil {
		t.Errorf("invalid key is not work")
	}

	invalidKey2 := tryable.Value("second")
	if invalidKey2 != nil {
		t.Errorf("invalid key is not work")
	}

}