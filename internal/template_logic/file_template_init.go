package template_logic

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Mallbrusss/project-generator/internal/models"
)

func CreateFileFromTemplate(dirPath string, file models.File, directory models.Directory) error {
	filePath := filepath.Join(dirPath, file.Name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("Use template for: %s", file.Name)

		tmplName := file.Name + ".tmpl"
		tmplPath := filepath.Join("templates", tmplName)

		if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
			log.Printf("No specific template found for %s. Using default.", file.Name)
			tmplPath = filepath.Join("templates", "default_template.go.tmpl")
		}

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			log.Printf("Error parsing template for %s: %v\n", file.Name, err)
			return err
		}

		f, err := os.Create(filePath)
		if err != nil {
			log.Printf("Can't create file %s: %v\n", file.Name, err)
			return err
		}
		defer f.Close()

		err = tmpl.Execute(f, map[string]interface{}{
			"PackageName": directory.Name,
		})
		if err != nil {
			log.Printf("Error executing template for %s: %v\n", file.Name, err)
			return err
		}

		log.Printf("File %s created with template\n", file.Name)
	} else {
		log.Printf("File %s already exists\n", file.Name)
	}
	return nil
}

func CreateEmptyFile(dirPath string, file models.File, directory models.Directory) {
	filePath := filepath.Join(dirPath, file.Name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		if err != nil {
			log.Printf("Can't create file %s: %v\n", file.Name, err)
			return
		}
		defer f.Close()

		if file.Name != "main.go" {
			_, err = f.WriteString(fmt.Sprint("package\t", directory.Name))
		} else {
			_, err = f.WriteString(fmt.Sprint("package\t", "main"))
		}

		log.Printf("File %s created without template_logic\n", file.Name)
	} else {
		log.Printf("File %s already exists\n", file.Name)
	}
}
