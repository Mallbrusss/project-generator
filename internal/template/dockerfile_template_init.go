package template

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func CreateDockerfile(basePath string) {
	dockerfilePath := filepath.Join(basePath, "Dockerfile")
	f, err := os.Create(dockerfilePath)
	if err != nil {
		log.Printf("Can't create Dockerfile: %v\n", err)
		return
	}
	defer f.Close()

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v\n", err)
	}

	// Строим путь к шаблону относительно рабочей директории
	tmplPath := filepath.Join(workingDir, "internal/template/dockerfile_template.go.tmpl")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Error parsing Dockerfile template: %v\n", err)
		return
	}

	err = tmpl.Execute(f, nil)
	if err != nil {
		log.Printf("Error executing Dockerfile template: %v\n", err)
		return
	}

	log.Println("Dockerfile created successfully")
}