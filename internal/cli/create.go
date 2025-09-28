package cli

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	cp "github.com/otiai10/copy"
)

func createDirectory(name string) error {
	if err := os.MkdirAll(name, 0755); err != nil {
		return fmt.Errorf("Unable to create directory %v: %v", name, err)
	}
	return nil
}

func copyTemplate(name string) error {
	configDirPath, err := getConfigDirPath()
	if err != nil {
		return fmt.Errorf("Not able to access config: %v", err)
	}
	template := "default.typ" // NOTE: Change this later
	templateSource := filepath.Join(configDirPath, "templates", template)

	if err := cp.Copy(templateSource, name+"/main.typ"); err != nil {
		return fmt.Errorf("Unable to copy template %v: %v", template, err)
	}

	return nil
}

func createToml(name string) (ProjectConfig, error) {
	var config ProjectConfig
	config, err := getProjectConfigFromGlobal()
	if err != nil {
		return config, err
	}

	if config.Title == "same" {
		config.Title = name
	}

	tomlPath := path.Join(name, "config.toml")
	file, err := os.Create(tomlPath)
	if err != nil {
		return config, fmt.Errorf("Unable to create project configuration: %v", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return config, fmt.Errorf("Unable to create project configuration: %v", err)
	}

	return config, nil
}

func create(name string) {
	if err := ensureConfigDirIsSetup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := createDirectory(name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err := createToml(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := copyTemplate(name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Document %v successfully created\n", name)
}
