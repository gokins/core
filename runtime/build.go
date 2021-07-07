package runtime

import "time"

type Build struct {
	Id         string                `yaml:"id"`
	PipelineId string                `yaml:"pipelineId"`
	Trigger    string                `yaml:"trigger"`
	Status     string                `yaml:"status"`
	Ref        string                `yaml:"ref"`
	Error      string                `yaml:"error"`
	Event      string                `yaml:"event"`
	Timestamp  time.Time             `yaml:"timestamp"`
	Title      string                `yaml:"title"`
	Message    string                `yaml:"message"`
	Vars       map[string]*Variables `yaml:"vars"`
	Started    time.Time             `yaml:"started"`
	Finished   time.Time             `yaml:"finished"`
	Created    time.Time             `yaml:"created"`
	Updated    time.Time             `yaml:"updated"`
	Before     string                `yaml:"before"`
	After      string                `yaml:"after"`
	Repo       *Repository           `yaml:"repo"`
	Stages     []*Stage              `yaml:"stages"`
}

type Variables struct {
	Name   string `yaml:"name,omitempty" json:"name"`
	Value  string `yaml:"value,omitempty" json:"value"`
	Secret bool   `yaml:"secret,omitempty" json:"secret"`
}

type Repository struct {
	Name     string `yaml:"name"`
	Token    string `yaml:"token"`
	Sha      string `yaml:"sha"`
	CloneURL string `yaml:"cloneUrl"`
}

type Stage struct {
	Id          string            `yaml:"id"`
	BuildId     string            `yaml:"buildId"`
	Name        string            `yaml:"name"`
	DisplayName string            `yaml:"displayName" `
	Stage       string            `yaml:"stage"`
	Status      string            `yaml:"status"`
	Event       string            `yaml:"event"`
	Error       string            `yaml:"error"`
	ExitCode    int               `yaml:"exitCode"`
	Started     time.Time         `yaml:"started"`
	Stopped     time.Time         `yaml:"stopped"`
	Finished    time.Time         `yaml:"finished"`
	Created     time.Time         `yaml:"created"`
	Updated     time.Time         `yaml:"updated"`
	Version     string            `yaml:"version"`
	OnSuccess   bool              `yaml:"onSuccess"`
	OnFailure   bool              `yaml:"onFailure"`
	Labels      map[string]string `yaml:"labels"`
	Steps       []*Step           `yaml:"steps"`
}

type Step struct {
	Id           string            `yaml:"id"`
	BuildId      string            `yaml:"buildId"`
	StageId      string            `yaml:"stageId"`
	Step         string            `yaml:"step"`
	DisplayName  string            `yaml:"displayName"`
	Name         string            `yaml:"name"`
	Env          map[string]string `yaml:"env"`
	Commands     interface{}       `yaml:"commands"`
	Status       string            `yaml:"status"`
	Event        string            `yaml:"event"`
	Error        string            `yaml:"error"`
	ExitCode     int               `yaml:"exitCode"`
	ErrIgnore    bool              `yaml:"errignore"`
	Started      time.Time         `yaml:"started"`
	Stopped      time.Time         `yaml:"stopped"`
	Finished     time.Time         `yaml:"finished"`
	Version      string            `yaml:"version"`
	Waits        []string          `yaml:"waits"`
	Image        string            `yaml:"image"`
	Artifacts    []*Artifact       `yaml:"artifacts"`
	UseArtifacts []*UseArtifact    `yaml:"useArtifacts"`
}

type Artifact struct {
	//Id        string    `yaml:"id"`
	BuildId string `yaml:"buildId"`
	StageId string `yaml:"stageId"`
	JobId   string `yaml:"stepId"`

	Scope      string `yaml:"scope"`      //archive,pipeline,env
	Repository string `yaml:"repository"` // archive,制品库ID
	Name       string `yaml:"name"`       //archive,pipeline,env
	Path       string `yaml:"path"`       //archive,pipeline

}

type UseArtifact struct {
	//Id        string    `yaml:"id"`
	BuildId string `yaml:"buildId"`
	StageId string `yaml:"stageId"`
	JobId   string `yaml:"stepId"`

	Scope      string `yaml:"scope"`      //archive,pipeline,env
	Repository string `yaml:"repository"` // archive,制品库ID
	Name       string `yaml:"name"`       //archive,pipeline,env
	IsForce    bool   `yaml:"isForce"`
	Path       string `yaml:"path"` //archive,pipeline

	SourceStage string `yaml:"sourceStage"` //pipeline
	SourceJob   string `yaml:"sourceStep"`  //pipeline

	Value string `yaml:"value"`
}
