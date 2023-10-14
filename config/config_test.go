package config

import "testing"

func TestConfig(t *testing.T) {
	config, err := ReadYAML("./../config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if config.Region == "" {
		t.Error("got empty REGION")
	}

	if config.AwsAccessKeyId == "" {
		t.Error("got empty AWS_ACCESS_KEY_ID")
	}

	if config.AwsSecretAccessKey == "" {
		t.Error("got empty AWS_SECRET_ACCESS_KEY")
	}

	if config.Endpoint != "" && config.Endpoint != LocalEndpoint {
		t.Errorf("got invalid ENDPOINT %s", config.Endpoint)
	}
}
