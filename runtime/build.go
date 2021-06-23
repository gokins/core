package runtime

import "time"

type Build struct {
	Id        string            `yaml:"id"`
	Trigger   string            `yaml:"trigger"`
	Status    string            `yaml:"status"`
	Ref       string            `yaml:"ref"`
	Error     string            `yaml:"error,omitempty"`
	Event     string            `yaml:"event"`
	Timestamp time.Time         `yaml:"timestamp"`
	Title     string            `yaml:"title,omitempty"`
	Message   string            `yaml:"message"`
	Params    map[string]string `yaml:"params,omitempty"`
	Started   time.Time         `yaml:"started"`
	Finished  time.Time         `yaml:"finished"`
	Created   time.Time         `yaml:"created"`
	Updated   time.Time         `yaml:"updated"`
	Before    string            `yaml:"before"`
	After     string            `yaml:"after"`
	Repo      *Repository       `yaml:"repo"`
	Variables map[string]string `yaml:"variables"`
	Stages    []*Stage          `yaml:"stages"`
}

type Repository struct {
	Name     string `yaml:"name"`
	Token    string `yaml:"token"`
	Sha      string `yaml:"sha"`
	CloneUrl string `yaml:"cloneUrl"`
}

type Stage struct {
	Id          string            `yaml:"id"`
	PipelineId  string            `yaml:"pipelineId"`
	BuildId     string            `yaml:"buildId"`
	Name        string            `yaml:"name"`
	DisplayName string            `yaml:"displayName" `
	Stage       string            `yaml:"stage"`
	Status      string            `yaml:"status"`
	Event       string            `yaml:"event"`
	Error       string            `yaml:"error"`
	ExitCode    int               `yaml:"exit_code"`
	Started     time.Time         `yaml:"started"`
	Stopped     time.Time         `yaml:"stopped"`
	Finished    time.Time         `yaml:"finished"`
	Created     time.Time         `yaml:"created"`
	Updated     time.Time         `yaml:"updated"`
	Version     string            `yaml:"version"`
	OnSuccess   bool              `yaml:"on_success"`
	OnFailure   bool              `yaml:"on_failure"`
	Labels      map[string]string `yaml:"labels"`
	Steps       []*Step           `yaml:"steps"`
}

type Step struct {
	Id              string            `yaml:"id"`
	StageId         string            `yaml:"stageId"`
	BuildId         string            `yaml:"buildId"`
	Step            string            `yaml:"step"`
	DisplayName     string            `yaml:"displayName"`
	Name            string            `yaml:"name"`
	Environments    map[string]string `yaml:"environments"`
	Commands        interface{}       `yaml:"commands"`
	Number          int               `yaml:"number"`
	Status          string            `yaml:"status"`
	Event           string            `yaml:"event"`
	Error           string            `yaml:"error,omitempty"`
	ErrIgnore       bool              `yaml:"errignore,omitempty"`
	ExitCode        int               `yaml:"exit_code"`
	Started         time.Time         `yaml:"started,omitempty"`
	Stopped         time.Time         `yaml:"stopped,omitempty"`
	Finished        time.Time         `yaml:"finished"`
	Version         string            `yaml:"version"`
	DependsOn       []string          `yaml:"dependsOn"`
	Image           string            `yaml:"image"`
	Artifacts       []*Artifact       `yaml:"artifacts"`
	DependArtifacts []*DependArtifact `yaml:"dependArtifacts"`
	IsClone         bool              `yaml:"isClone"`
}

type Artifact struct {
	Id        string    `yaml:"id"`
	BuildId   string    `yaml:"buildId"`
	StageId   string    `yaml:"stageId"`
	JobId     string    `yaml:"stepId"`
	BuildName string    `yaml:"buildName"`
	StageName string    `yaml:"stageName"`
	JobName   string    `yaml:"stepName"`
	Status    string    `yaml:"status"`
	Created   time.Time `yaml:"created"`
	Updated   time.Time `yaml:"updated"`

	Name  string `yaml:"name"`
	Scope string `yaml:"scope"`
	Path  string `yaml:"path"`

	Artifactory string `yaml:"artifactory"`
	Repository  string `yaml:"repository"`

	Value string `yaml:"value"`
}

type DependArtifact struct {
	Id        string    `yaml:"id"`
	BuildId   string    `yaml:"buildId"`
	StageId   string    `yaml:"stageId"`
	JobId     string    `yaml:"stepId"`
	BuildName string    `yaml:"buildName"`
	StageName string    `yaml:"stageName"`
	JobName   string    `yaml:"stepName"`
	Status    string    `yaml:"status"`
	Created   time.Time `yaml:"created"`
	Updated   time.Time `yaml:"updated"`

	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	IsForce     bool   `yaml:"isForce"`
	Artifactory string `yaml:"artifactory"`
	Repository  string `yaml:"repository"`
	Target      string `yaml:"target"`

	SourceStage string `yaml:"sourceStage"`
	SourceJob   string `yaml:"sourceStep"`

	Value string `yaml:"value"`
}
