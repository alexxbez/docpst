package cli

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/alexxbez/docpst/internal/assets"
)

type ProjectConfig struct {
	Title           string `toml:"title"`
	Author          string `toml:"author"`
	defaultTemplate string `toml:"dafult_template"`
	Source          string `toml:"source"`
	Dest            string `toml:"dest"`
}

func ensureConfigDirIsSetup() error {
	configDirPath, err := getConfigDirPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(configDirPath); err != nil {
		if os.IsNotExist(err) {
			if setupErr := setupConfigDir(configDirPath); setupErr != nil {
				return setupErr
			}
		} else {
			return err
		}
	}
	return nil
}

func getConfigDirPath() (string, error) {
	configDirPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDirPath = path.Join(configDirPath, "docpst")

	return configDirPath, nil
}

func setupConfigDir(configDirPath string) error {
	if err := os.MkdirAll(configDirPath, 0755); err != nil {
		return fmt.Errorf("Unable to create %v directory: %v", configDirPath, err)
	}

	templatesPath := path.Join(configDirPath, "templates")
	if err := os.MkdirAll(templatesPath, 0755); err != nil {
		return fmt.Errorf("Unable to create templates directory: %v", err)
	}

	defaultTemplateName := "default.typ"
	if err := copyEmbeddedFile(assets.Assets, defaultTemplateName, path.Join(templatesPath, defaultTemplateName)); err != nil {
		return err
	}

	tomlConfigName := "config.toml"
	if err := copyEmbeddedFile(assets.Assets, tomlConfigName, path.Join(configDirPath, tomlConfigName)); err != nil {
		return err
	}

	return nil
}

func copyEmbeddedFile(embeddedFiles fs.FS, fileName string, destPath string) error {
	srcFile, err := embeddedFiles.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed to open embedded file %v: %v", fileName, err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("Failed to create destination file %v: %v", destFile, err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("Failed to copy content: %v", err)
	}

	return nil
}

func getProjectConfigFromGlobal() (ProjectConfig, error) {
	var config struct {
		Project ProjectConfig `toml:"project"`
	}

	configPath, err := getConfigDirPath()
	if err != nil {
		return config.Project, err
	}
	tomlPath := path.Join(configPath, "config.toml")

	if _, err := toml.DecodeFile(tomlPath, &config); err != nil {
		return config.Project, fmt.Errorf("Unable to read configuration: %v", err)
	}

	return config.Project, err
}
