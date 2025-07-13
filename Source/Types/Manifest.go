package Types

type Manifest struct {
	Application struct {
		ID         string `yaml:"ID"`
		Name       string `yaml:"name"`
		Icon       string `yaml:"icon"`
		Entrypoint string `yaml:"entrypoint"`
	} `yaml:"Application"`

	Store struct {
		Author struct {
			ID   string `yaml:"ID"`
			Type string `yaml:"type"`
		} `yaml:"author"`
		Description string `yaml:"description"`
		Versions    []struct {
			Version     string `yaml:"version"`
			Date        string `yaml:"date"`
			Private     bool   `yaml:"private"`
			Description string `yaml:"description"`
		} `yaml:"versions"`
		Gallery []struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
			File string `yaml:"file"`
		} `yaml:"gallery"`
	} `yaml:"Store"`

	Permissions []struct {
		ID         string `yaml:"ID"`
		Permission string `yaml:"permission"`
		Reason     string `yaml:"reason"`
		Data       any    `yaml:"data,omitempty"`
	} `yaml:"Permissions"`
}

type Permission struct {
	ID         string      `yaml:"ID"`
	Permission string      `yaml:"permission"`
	Reason     string      `yaml:"reason"`
	Data       interface{} `yaml:"data,omitempty"`
}
