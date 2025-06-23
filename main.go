package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	directorylogic "github.com/Mallbrusss/project-generator/internal/directory_logic"
	parselogic "github.com/Mallbrusss/project-generator/internal/parse_logic"
	"github.com/Mallbrusss/project-generator/internal/template_logic"
)

func main() {
	helpL := flag.Bool("help", false, "Print help")
	helpS := flag.Bool("h", false, "Print help")
	output := flag.String("output", "output", "Output directory")
	configFile := flag.String("config-file", "", "Use -config-file to specify the path to the configuration file")

	flag.Parse()

	// help logic block
	if *helpL || *helpS {
		if *helpL && *helpS {
			fmt.Println("Error: don't use '-h' & '-help' together")
			os.Exit(1)
		}

		// Печать помощи
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *configFile == "" {
		fmt.Println("Error: config file is required")
		fmt.Println("Usage: project-generator -config-file=<path_to_config>")
		os.Exit(1)
	}

	project, err := parselogic.ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	projectRoot := filepath.Join(*output, project.ProjectName)

	for _, directory := range project.Directories {
		directorylogic.CreateDirectoryStructure(projectRoot, directory)
	}

	//allFeatures := []models.Feature{
	//	&pgx.PgxFeature{Enable: project.Contains("pgx")},
	//}
	//err = features.GenerateFeature(projectRoot, allFeatures)
	if project.Docker {
		template_logic.CreateDockerfile(projectRoot)
	}

	initializeGoProj(projectRoot, project.ProjectName)

}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Error executing command '%s %v': %v\n", command, args, err)
	} else {
		log.Printf("Command '%s %v' executed successfully\n", command, args)
	}
}

func initializeGoProj(basePath, projectName string) {
	err := os.Chdir(basePath)
	if err != nil {
		log.Printf("Error changing directory to %s: %v\n", basePath, err)
		return
	}

	runCommand("go", "mod", "init", projectName)
	runCommand("go", "mod", "tidy")

	//TODO: Чекать зависимости в ямле
	// runCommand("go", "get", "github.com/labstack/echo/v4")
}
