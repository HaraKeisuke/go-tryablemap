package tryablemap

type TryableType int

const (
	Map TryableType = iota
	Array
)

type TryableMap struct {
	Map *map[string]interface{}
}

type TryableArray struct {
	Map   *map[string][]interface{}
	Array []interface{}
}

func NewTryableMap(_map map[string]interface{}) *TryableMap {
	return &TryableMap{
		Map: &_map,
	}
}

func NewTryableArray(_map map[string][]interface{}) *TryableArray {
	return &TryableArray{
		Map: &_map,
	}
}

func (tmap *TryableMap) Try(key string) *TryableMap {
	if tmap.Map == nil {
		return &TryableMap{}
	}

	_val, ok := (*tmap.Map)[key].(map[string]interface{})
	if ok {
		return &TryableMap{
			Map: &_val,
		}
	}
	return &TryableMap{}
}

func (tmap *TryableArray) Try(key string) *TryableArray {
	if tmap.Map == nil {
		return &TryableArray{}
	}
	_val, ok := (*tmap.Map)[key]
	if ok {
		return &TryableArray{
			Array: _val,
		}
	}
	return &TryableArray{}
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

	maxLen := len(tmap.Array)
	if maxLen < key {
		return nil
	}

	_val := tmap.Array[key]
	return _val
}
