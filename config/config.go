package config

import (
	"os"

	// "github.com/cloudfoundry-community/go-credhub"
	"gopkg.in/yaml.v3"
)

const ConfigEnvVarName string = "CONFIG_SERVER_BROKER_CONFIG"

type Config struct {
	CredhubURL          string `yaml:"credhub_url"`
	Auth                Auth   `yaml:"broker_auth"`
	ServiceName         string `yaml:"service_name"`
	ServiceID           string `yaml:"service_id"`
	BasicPlanId         string `yaml:"basic_plan_id"`
	BasicPlanName       string `yaml:"basic_plan_name"`
	Description         string `yaml:"description"`
	LongDescription     string `yaml:"long_description"`
	ProviderDisplayName string `yaml:"provider_display_name"`
	DocumentationURL    string `yaml:"documentation_url"`
	SupportURL          string `yaml:"support_url"`
	DisplayName         string `yaml:"display_name"`
	IconImage           string `yaml:"icon_image"`
}

type Auth struct {
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

func ParseConfig() (Config, error) {
	configJson := os.Getenv(ConfigEnvVarName)
	if configJson == "" {
		panic(ConfigEnvVarName + " not set")
	}
	var config Config

	return config, yaml.Unmarshal([]byte(configJson), &config)
}
