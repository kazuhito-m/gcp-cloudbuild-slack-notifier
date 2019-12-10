package cloudbuild

import "time"

const STATUS_SUCCESS = "SUCCESS"
const STATUS_FAILURE = "FAILURE"
const STATUS_INTERNAL_ERROR = "INTERNAL_ERROR"
const STATUS_TIMEOUT = "TIMEOUT"

const STATUS_QUEUED = "QUEUED"
const STATUS_WORKING = "WORKING"

func Statuses() []string {
	return []string{
		STATUS_SUCCESS,
		STATUS_FAILURE,
		STATUS_INTERNAL_ERROR,
		STATUS_TIMEOUT,
		STATUS_QUEUED,
		STATUS_WORKING,
	}
}

func InCloudBuildStatus(status string) bool {
	for _, i := range Statuses() {
		if i == status {
			return true
		}
	}
	return false
}

type CloudBuildResult struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Status    string `json:"status"`
	Source    struct {
		RepoSource struct {
			ProjectID  string `json:"projectId"`
			RepoName   string `json:"repoName"`
			BranchName string `json:"branchName"`
		} `json:"repoSource"`
	} `json:"source"`
	Steps []struct {
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
	} `json:"steps"`
	Results struct {
		BuildStepImages []string `json:"buildStepImages"`
	} `json:"results"`
	CreateTime       time.Time `json:"createTime"`
	StartTime        time.Time `json:"startTime"`
	FinishTime       time.Time `json:"finishTime"`
	Timeout          string    `json:"timeout"`
	LogsBucket       string    `json:"logsBucket"`
	SourceProvenance struct {
		ResolvedRepoSource struct {
			ProjectID string `json:"projectId"`
			RepoName  string `json:"repoName"`
			CommitSha string `json:"commitSha"`
		} `json:"resolvedRepoSource"`
	} `json:"sourceProvenance"`
	BuildTriggerID string `json:"buildTriggerId"`
	Options        struct {
		SubstitutionOption string `json:"substitutionOption"`
		Logging            string `json:"logging"`
	} `json:"options"`
	LogURL string   `json:"logUrl"`
	Tags   []string `json:"tags"`
	Timing struct {
		BUILD struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"BUILD"`
		FETCHSOURCE struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"FETCHSOURCE"`
	} `json:"timing"`
}

func (i CloudBuildResult) Ok() bool {
	return i.Status == "SUCCESS"
}

func (i CloudBuildResult) Ng() bool {
	return i.Status == "FAILURE"
}
