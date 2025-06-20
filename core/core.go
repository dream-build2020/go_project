package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func SetConfig() (map[string]string, error) {
	var config AppConfig

	data, err := os.ReadFile("config/coreConfig.yaml")
	if err != nil {
		return nil, fmt.Errorf("read dir config.yaml failed: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config.yaml failed: %v", err)
	}

	username := config.Account.Username
	password := config.Account.Password
	result := make(map[string]string)
	result["username"] = username
	result["password"] = password

	return result, err
}

func SetMap(mapping map[string]string) {
	for _, value := range mapping {
		fmt.Println(value)
	}
}
