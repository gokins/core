package runtime

import "time"

type BuildShow struct {
	Id         string       `json:"id"`
	PipelineId string       `json:"pipelineId"`
	Status     string       `json:"status"`
	Error      string       `json:"error"`
	Event      string       `json:"event"`
	Started    time.Time    `json:"started"`
	Finished   time.Time    `json:"finished"`
	Created    time.Time    `json:"created"`
	Updated    time.Time    `json:"updated"`
	Stages     []*StageShow `json:"stages"`
}

type StageShow struct {
	Id       string      `json:"id"`
	BuildId  string      `json:"buildId"`
	Status   string      `json:"status"`
	Event    string      `json:"event"`
	Error    string      `json:"error"`
	Started  time.Time   `json:"started"`
	Stopped  time.Time   `json:"stopped"`
	Finished time.Time   `json:"finished"`
	Created  time.Time   `json:"created"`
	Updated  time.Time   `json:"updated"`
	Steps    []*StepShow `json:"steps"`
}
type StepShow struct {
	Id       string    `json:"id"`
	BuildId  string    `json:"buildId"`
	StageId  string    `json:"stageId"`
	Status   string    `json:"status"`
	Event    string    `json:"event"`
	Error    string    `json:"error"`
	ExitCode int       `json:"exitCode"`
	Started  time.Time `json:"started"`
	Stopped  time.Time `json:"stopped"`
	Finished time.Time `json:"finished"`
}
type CmdShow struct {
	Id       string    `json:"id"`
	BuildId  string    `json:"buildId"`
	StepId   string    `json:"stageId"`
	Status   string    `json:"status"`
	Started  time.Time `json:"started"`
	Finished time.Time `json:"finished"`
}
