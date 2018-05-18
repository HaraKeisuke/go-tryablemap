package tryablemap

type TryableMap struct {
	Map map[string]interface{}
}

func NewTryableMap(_map map[string]interface{}) *TryableMap {
	return &TryableMap{
		Map: _map,
	}
}

func (tmap *TryableMap) Try(key string) *TryableMap {
	if tmap.Map == nil {
		return nil
	}

	if val, ok := tmap.Map[key]; ok {
		_val := val.(map[string]interface{})
		return &TryableMap{
			Map: _val,
		}
	} else {
		return nil
	}
}
