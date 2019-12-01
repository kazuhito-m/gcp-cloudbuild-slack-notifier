package cloudbuild

import (
	"encoding/json"
	"log"
)

func Parse(jsonText string) CluodBuildResult {
	bytes := []byte(jsonText)
	var decode_data CluodBuildResult
	if err := json.Unmarshal(bytes, &decode_data); err != nil {
		log.Fatal(err)
	}
	return decode_data
}
