package tryablemap

type TryableMap struct {
	Map *map[string]interface{}
}

func NewTryableMap(_map map[string]interface{}) *TryableMap {
	return &TryableMap{
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
