package healthcheck

type Response struct {
	App App `json:"app"`
}

type App struct {
	Codebase    *Codebase    `json:"codebase,omitempty"`
	Environment *Environment `json:"environment,omitempty"`
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	GoVersion   string       `json:"go"`
}

type Codebase struct {
	Branch     string `json:"branch,omitempty"`
	CommitHash string `json:"commit,omitempty"`
}

type Environment struct {
	Name       string `json:"name,omitempty"`
	Region     string `json:"region,omitempty"`
	InstanceId string `json:"instanceId,omitempty"`
}
