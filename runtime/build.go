package runtime

import "time"

type Build struct {
	Id        string            `json:"id"`
	Trigger   string            `json:"trigger"`
	Status    string            `json:"status"`
	Ref       string            `json:"ref"`
	Error     string            `json:"error,omitempty"`
	Event     string            `json:"event"`
	Timestamp time.Time         `json:"timestamp"`
	Title     string            `json:"title,omitempty"`
	Message   string            `json:"message"`
	Params    map[string]string `json:"params,omitempty"`
	Started   time.Time         `json:"started"`
	Finished  time.Time         `json:"finished"`
	Created   time.Time         `json:"created"`
	Updated   time.Time         `json:"updated"`
	Before    string            `json:"before"`
	After     string            `json:"after"`
	Stages    []*Stage          `json:"stages"`
}

type Stage struct {
	Id          string            `json:"id"`
	PipelineId  string            `json:"pipelineId"`
	BuildId     string            `json:"buildId"`
	Name        string            `json:"name"`
	DisplayName string            `json:"displayName" `
	Stage       string            `json:"stage"`
	Status      string            `json:"status"`
	Event       string            `json:"event"`
	Error       string            `json:"error"`
	ExitCode    int               `json:"exit_code"`
	Started     time.Time         `json:"started"`
	Stopped     time.Time         `json:"stopped"`
	Finished    time.Time         `json:"finished"`
	Created     time.Time         `json:"created"`
	Updated     time.Time         `json:"updated"`
	Version     string            `json:"version"`
	OnSuccess   bool              `json:"on_success"`
	OnFailure   bool              `json:"on_failure"`
	Labels      map[string]string `json:"labels"`
	Steps       []*Step           `yaml:"steps"`
}

type Step struct {
	Id              string            `json:"id"`
	StageId         string            `json:"stageId"`
	BuildId         string            `json:"buildId"`
	Step            string            `json:"step"`
	DisplayName     string            `json:"displayName"`
	Name            string            `son:"name"`
	Environments    map[string]string `json:"environments"`
	Commands        []interface{}     `yaml:"commands" json:"-"`
	Command         string            `json:"'command'"`
	Number          int               `json:"number"`
	Status          string            `json:"status"`
	Event           string            `json:"event"`
	Error           string            `json:"error,omitempty"`
	ErrIgnore       bool              `json:"errignore,omitempty"`
	ExitCode        int               `json:"exit_code"`
	Started         time.Time         `json:"started,omitempty"`
	Stopped         time.Time         `json:"stopped,omitempty"`
	Finished        time.Time         `json:"finished"`
	Version         string            `json:"version"`
	DependsOn       []string          `json:"dependsOn"`
	Image           string            `json:"image"`
	Artifacts       []*Artifact       `json:"artifacts"`
	DependArtifacts []*DependArtifact `json:"dependArtifacts"`
}

type Artifact struct {
	Id        string
	BuildId   string    `json:"buildId"`
	StageId   string    `json:"stageId"`
	JobId     string    `json:"stepId"`
	BuildName string    `json:"buildName"`
	StageName string    `json:"stageName"`
	JobName   string    `json:"stepName"`
	Status    string    `json:"status"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`

	Name  string `json:"name"`
	Scope string `json:"scope"`
	Path  string `json:"path"`

	Artifactory string `json:"artifactory"`
	Repository  string `json:"repository"`

	Value string `json:"value"`
}

type DependArtifact struct {
	Id        string
	BuildId   string    `json:"buildId"`
	StageId   string    `json:"stageId"`
	JobId     string    `json:"stepId"`
	BuildName string    `json:"buildName"`
	StageName string    `json:"stageName"`
	JobName   string    `json:"stepName"`
	Status    string    `json:"status"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`

	Name        string `json:"name"`
	Type        string `json:"type"`
	IsForce     bool   `json:"isForce"`
	Artifactory string `json:"artifactory"`
	Repository  string `json:"repository"`
	Target      string `json:"target"`

	SourceStage string `json:"sourceStage"`
	SourceJob   string `json:"sourceStep"`

	Value string `json:"value"`
}
