package runtime

import "time"

type Build struct {
	Id                string
	PipelineId        string
	PipelineVersionId string
	Trigger           string
	Status            string
	Ref               string
	Error             string
	Event             string
	Timestamp         time.Time
	Title             string
	Message           string
	Vars              map[string]*Variables
	Started           time.Time
	Finished          time.Time
	Created           time.Time
	Updated           time.Time
	Before            string
	After             string
	Repo              *Repository
	Stages            []*Stage
}

type Variables struct {
	Name   string
	Value  string
	Secret bool
}

type Repository struct {
	Name     string
	Token    string
	Sha      string
	CloneURL string
}

type Stage struct {
	Id          string
	BuildId     string
	Name        string
	DisplayName string
	Stage       string
	Status      string
	Event       string
	Error       string
	ExitCode    int
	Started     time.Time
	Stopped     time.Time
	Finished    time.Time
	Created     time.Time
	Updated     time.Time
	Version     string
	OnSuccess   bool
	OnFailure   bool
	Labels      map[string]string
	Steps       []*Step
}

type Step struct {
	Id           string
	BuildId      string
	StageId      string
	Step         string
	DisplayName  string
	Name         string
	Input        map[string]string
	Env          map[string]string
	Commands     interface{}
	Status       string
	Event        string
	Error        string
	ExitCode     int
	ErrIgnore    bool
	Started      time.Time
	Stopped      time.Time
	Finished     time.Time
	Version      string
	Waits        []string
	Image        string
	Artifacts    []*Artifact
	UseArtifacts []*UseArtifact
}

type Artifact struct {
	Scope      string //archive,pipeline,env
	Repository string // archive,制品库ID
	Name       string //archive,pipeline,env
	Path       string //archive,pipeline
}

type UseArtifact struct {
	Scope      string //archive,pipeline,env
	Repository string // archive,制品库ID
	Name       string //archive,pipeline,env
	Path       string //archive,pipeline
	IsUrl      bool
	Alias      string
	//IsForce    bool

	SourceStage string //pipeline
	SourceStep  string //pipeline
}
