package tryablemap

import (
	"reflect"
)

func (tmap *TryableMap) Try(key string) *TryableMap {
	if tmap.Map == nil {
		return &TryableMap{}
	}

	val := (*tmap.Map)[key]
	if val == nil {
		return &TryableMap{}
	}

	hash := makeMap(val)
	if _val, ok := (*hash).(map[string]interface{}); ok {
		return &TryableMap{
			Map: &_val,
		}
	}
	return &TryableMap{}
}

func (tmap *TryableMap) TryArray(key string) *TryableArray {
	if tmap.Map == nil {
		return &TryableArray{}
	}

	val := (*tmap.Map)[key]
	if val == nil {
		return &TryableArray{}
	}

	array := makeArray(val)

	return &TryableArray{
		Array: &array,
	}
	return &TryableArray{}
}

func (tmap *TryableArray) Try(key int) *TryableMap {
	if tmap.Array == nil {
		return &TryableMap{}
	}
	array := makeArray(*tmap.Array)
	if _val, ok := array[key].(map[string]interface{}); ok {
		return &TryableMap{
			Map: &_val,
		}
	}
	return &TryableMap{}
}

func (tmap *TryableArray) TryArray(key int) *TryableArray {
	if tmap.Array == nil {
		return &TryableArray{}
	}

	maxLen := len(*tmap.Array)
	if maxLen <= key {
		return &TryableArray{}
	}

	val := (*tmap.Array)[key]
	if val == nil {
		return &TryableArray{}
	}

	array := makeArray(val)
	return &TryableArray{
		Array: &array,
	}
}

func (tmap *TryableMap) Value(key string) interface{} {
	if tmap.Map != nil {
		if val, ok := (*tmap.Map)[key]; ok {
			return val
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (tmap *TryableArray) Value(key int) interface{} {
	if tmap.Array == nil {
		return nil
	}
	maxLen := len(*tmap.Array)
	if maxLen <= key {
		return nil
	}

	_val := (*tmap.Array)[key]
	return _val
}

func makeMap(object interface{}) *interface{} {
	t := reflect.TypeOf(object)
	if t.Kind() == reflect.Map {
		v := reflect.ValueOf(object)
		it := reflect.TypeOf((*interface{})(nil)).Elem()
		m := reflect.MakeMap(reflect.MapOf(t.Key(), it))
		for _, mk := range v.MapKeys() {
			m.SetMapIndex(mk, v.MapIndex(mk))
		}
		result := m.Interface()
		return &result
	} else {
		return nil
	}
}

func makeArray(object interface{}) []interface{} {
	t := reflect.TypeOf(object)
	if t.Kind() == reflect.Slice {
		v := reflect.ValueOf(object)
		array := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			array[i] = v.Index(i).Interface()
		}
		return array
	}
	return []interface{}{}
}
