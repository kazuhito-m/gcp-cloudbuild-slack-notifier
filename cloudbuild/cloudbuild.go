package cloudbuild

import (
	"encoding/json"
	"log"
)

func Parse(bytes []byte) CloudBuildResult {
	var decode_data CloudBuildResult
	if err := json.Unmarshal(bytes, &decode_data); err != nil {
		log.Fatal(err)
	}
	return decode_data
}
