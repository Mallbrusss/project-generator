package models

type Project struct {
	ProjectName string      `yaml:"project_name"`
	Directories []Directory `yaml:"directories"`
	Docker      bool        `yaml:"dockerfile,omitempty"`
	Features    []string
}
type Directory struct {
	Name           string      `yaml:"name"`
	Description    string      `yaml:"description"`
	Subdirectories []Directory `yaml:"subdirectories,omitempty"`
	Files          []File      `yaml:"files,omitempty"`
	IsCmdDefault   bool        `yaml:"is_cmd_default,omitempty"`
}
type File struct {
	Name        string `yaml:"name"`
	Default     bool   `yaml:"default"`
	Description string
}

func (project *Project) Contains(feature string) bool {
	for _, featureName := range project.Features {
		if featureName == feature {
			return true
		}
	}
	return false
}
