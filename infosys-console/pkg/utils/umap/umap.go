package umap

import (
	"encoding/json"
)

func StructToMapByJsonTag(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}

func StructJsonNames(obj interface{}) []string {

	m, err := StructToMapByJsonTag(obj)
	if err != nil {
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
