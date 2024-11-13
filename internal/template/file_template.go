package template

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/Mallbrusss/project-generator/internal/models"
)

func CreateFileFromTemplate(dirPath string, file models.File, directory models.Directory) {
	filePath := filepath.Join(dirPath, file.Name)

	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Получаем шаблон для этого файла
		// templateStr, exists := templates[file.Name]
		// if !exists {
		// 	// Если специфического шаблона нет, используем дефолтный шаблон
		// 	log.Printf("No template found for %s, using default template\n", file.Name)
		templateStr := defaultTemplate
		// }

		// Парсим шаблон
		tmpl, err := template.New(file.Name).Parse(templateStr)
		if err != nil {
			log.Printf("Error parsing template for %s: %v\n", file.Name, err)
			return
		}

		// Открываем файл для записи
		f, err := os.Create(filePath)
		if err != nil {
			log.Printf("Can't create file %s: %v\n", file.Name, err)
			return
		}
		defer f.Close()

		// Заполняем шаблон и записываем в файл
		err = tmpl.Execute(f, map[string]interface{}{
			"PackageName": directory.Name,
			"Description": file.Description,
			"FuncName":    file.Name[:len(file.Name)-3],
		})
		if err != nil {
			log.Printf("Error executing template for %s: %v\n", file.Name, err)
			return
		}

		log.Printf("File %s created with template\n", file.Name)
	} else {
		log.Printf("File %s already exists\n", file.Name)
	}
}
