package directorylogic

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Mallbrusss/project-generator/internal/models"
	"github.com/Mallbrusss/project-generator/internal/template"
)

func CreateDirectoryStructure(basePath string, directory models.Directory) {
	dirPath := filepath.Join(basePath, directory.Name)

	err := os.MkdirAll(dirPath, os.ModePerm)

	if err != nil {
		log.Printf("Can`t create folder %s: %v\n", directory.Name, err)
	} else {
		log.Printf("folder %s created\n", directory.Name)
	}

	for _, subdir := range directory.Subdirectories {
		CreateDirectoryStructure(dirPath, subdir)

	}

	for _, file := range directory.Files {
		template.CreateFileFromTemplate(dirPath, file, directory)
	}

	

}
