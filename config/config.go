package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type AWS struct {
	Region             string `yaml:"REGION"`
	AwsAccessKeyId     string `yaml:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `yaml:"AWS_SECRET_ACCESS_KEY"`
	Endpoint           string `yaml:"ENDPOINT"`
}

func ReadYAML(path string) (*AWS, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	awsConfig := &AWS{}
	err = yaml.Unmarshal(file, awsConfig)
	if err != nil {
		return nil, err
	}

	return awsConfig, nil
}
