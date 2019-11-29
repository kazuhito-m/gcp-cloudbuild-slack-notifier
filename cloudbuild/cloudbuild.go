package cloudbuild

import (
	"encoding/json"
	"log"
)

func Parse(json_text string) CluodBuildResult {
	bytes := []byte(json_text)
	var decode_data CluodBuildResult
	if err := json.Unmarshal(bytes, &decode_data); err != nil {
		log.Fatal(err)
	}
	return decode_data
}
