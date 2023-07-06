package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token    string   `yaml:"token"`
	Policies []Policy `yaml:"policies"`
}

type Policy struct {
	Groups      []string `yaml:"groups"`
	Recursive   bool     `yaml:"recursive" default:"false"`
	MergeMethod string   `yaml:"merge_method"`
}

func ReadConfig(filename string) (Config, error) {
	// Read the YAML config file
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the config file into Config struct
	var config Config
	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatal(err)
	}

	// If the token is not set on the config file, try to get the token from GITLAB_TOKEN environment
	if config.Token == "" {
		token := os.Getenv("GITLAB_TOKEN")
		config.Token = token
	}

	// If the is not set by the config file or the environment, raise an error
	if config.Token == "" {
		return config, errors.New("token not provided")
	}

	return config, nil
}
