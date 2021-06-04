package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gleich/logoru"
	"gopkg.in/yaml.v3"
)

// Configuration outline
type Outline struct {
	Levels []string `yaml:"levels"`
}

// Get the configuration folder name
func GetFolderName() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logoru.Critical("Failed to get homedir for config;", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, "contrihat-config")
}

// Read from the config.yml file
func ReadGeneral() Outline {
	// Reading the binary from the file
	data, err := ioutil.ReadFile(filepath.Join(GetFolderName(), "config.yml"))
	if err != nil {
		logoru.Warning("Failed to read from configuration file;", err)
		return Outline{}
	}

	// Parsing yaml
	var outline Outline
	err = yaml.Unmarshal(data, &outline)
	if err != nil {
		logoru.Critical("Failed to parse configuration file")
		os.Exit(1)
	}
	return outline
}
