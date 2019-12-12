package cloudbuild

import "time"

type CloudBuildStep struct {
	Name       string   `json:"name"`
	Args       []string `json:"args"`
	ID         string   `json:"id,omitempty"`
	Entrypoint string   `json:"entrypoint"`
	Timing     struct {
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
	} `json:"timing"`
	PullTiming struct {
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
	} `json:"pullTiming"`
	Status  string   `json:"status"`
	WaitFor []string `json:"waitFor,omitempty"`
}

func (i CloudBuildStep) Description() string {
	if i.ID == "" {
		return i.Name
	}
	return i.ID + "(" + i.Name + ")"
}
