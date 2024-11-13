package parselogic

import (
	"log"
	"os"

	"github.com/Mallbrusss/project-generator/internal/models"
	"gopkg.in/yaml.v3"
)

func ReadConfig(filepath string) (models.Project, error) {
	var project models.Project
	data, err := os.ReadFile(filepath)
	if err != nil{
		log.Printf("error while reading conf file: %v", err)
		return project, err
	}

	err = yaml.Unmarshal(data, &project)
	if err != nil{
		log.Printf("error reading configuration file: %v", err)
		return project, err
	}

	return project, nil
}
