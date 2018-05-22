package tryablemap

type TryableMap struct {
	Map *map[string]interface{}
}

type TryableArray struct {
	Array *[]interface{}
}

func NewTryableMap(_map interface{}) *TryableMap {
	hash := makeMap(_map)
	if val, ok := (*hash).(map[string]interface{}); ok {
		return &TryableMap{
			Map: &val,
		}
	}
	return &TryableMap{}
}

func NewTryableArray(_array interface{}) *TryableArray {
	array := makeArray(_array)
	return &TryableArray{
		Array: &array,
	}
}

func New(_map interface{}) *TryableMap {
	return NewTryableMap(_map)
}

func NewArray(_array interface{}) *TryableArray {
	return NewTryableArray(_array)
}
