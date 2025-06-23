package models

type Feature interface {
	Name() string
	Enabled() bool
	Generate(outputDir string) error
}
