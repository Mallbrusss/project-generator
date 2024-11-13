package models

type Project struct {
	ProjectName string      `yaml:"project_name"`
	Directories []Directory `yaml:"directories"`
}

type Directory struct {
	Name           string      `yaml:"name"`
	Description    string      `yaml:"description"`
	Subdirectories []Directory `yaml:"subdirectories,omitempty"`
	Files          []File      `yaml:"files,omitempty"`
}

type File struct {
	Name        string `yaml:"name"`
	Template    string
	Description string
}
