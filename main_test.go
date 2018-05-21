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

func TestTryableArray(t *testing.T) {
	data := map[string][]interface{}{
		"sample": []interface{}{
			"first", "second", "third",
		},
	}
	tryable := tryablemap.NewTryableArray(data)
	first := tryable.Try("sample").Value(0)
	if first != "first" {
		t.Errorf("Second is not invalid")
	}

	second := tryable.Try("sample").Value(1)
	if second != "second" {
		t.Errorf("Second is not invalid")
	}

	outOfRange := tryable.Try("sample").Value(5)
	if outOfRange != nil {
		t.Errorf("OutOfRange is danger. It may happen panic.")
	}
}
