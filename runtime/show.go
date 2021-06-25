package runtime

import "time"

type BuildShow struct {
	Id         string       `yaml:"id"`
	PipelineId string       `yaml:"pipelineId"`
	Status     string       `yaml:"status"`
	Error      string       `yaml:"error,omitempty"`
	Event      string       `yaml:"event"`
	Started    time.Time    `yaml:"started"`
	Finished   time.Time    `yaml:"finished"`
	Created    time.Time    `yaml:"created"`
	Updated    time.Time    `yaml:"updated"`
	Stages     []*StageShow `yaml:"stages"`
}

type StageShow struct {
	Id       string      `yaml:"id"`
	BuildId  string      `yaml:"buildId"`
	Status   string      `yaml:"status"`
	Event    string      `yaml:"event"`
	Error    string      `yaml:"error"`
	Started  time.Time   `yaml:"started"`
	Stopped  time.Time   `yaml:"stopped"`
	Finished time.Time   `yaml:"finished"`
	Created  time.Time   `yaml:"created"`
	Updated  time.Time   `yaml:"updated"`
	Steps    []*StepShow `yaml:"steps"`
}
type StepShow struct {
	Id       string    `yaml:"id"`
	StageId  string    `yaml:"stageId"`
	BuildId  string    `yaml:"buildId"`
	Status   string    `yaml:"status"`
	Event    string    `yaml:"event"`
	Error    string    `yaml:"error,omitempty"`
	ExitCode int       `yaml:"exit_code"`
	Started  time.Time `yaml:"started,omitempty"`
	Stopped  time.Time `yaml:"stopped,omitempty"`
	Finished time.Time `yaml:"finished"`
}
