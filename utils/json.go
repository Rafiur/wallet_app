package utils

import "encoding/json"

func JsonCast(src, target interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, target)
}
