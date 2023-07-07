package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const GitlabTokenEnvkey = "GITLAB_TOKEN"

type Config struct {
	App      AppConfig `yaml:"app"`
	Policies []Policy  `yaml:"policies"`
}

type AppConfig struct {
	Token string `yaml:"token"`
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

	// If GitLab token is set by GITLAB_TOKEN environment, override the token provided by the configuration file
	token := os.Getenv(GitlabTokenEnvkey)
	if token != "" {
		config.App.Token = token
	}

	// If the GitLab token is not set by the config file or the environment, return an empty config and an error
	if config.App.Token == "" {
		return Config{}, errors.New("token not provided")
	}

	return config, nil
}
