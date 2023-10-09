package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbHost   string `json:"db_host"`
	DbPort   string `json:"db_port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func GetConfigs() (*Configs, error) {
	bytes, err := os.ReadFile("./internal/configs/configs.json")
	if err != nil {
		return nil, err
	}
	var newConfigs Configs

	err = json.Unmarshal(bytes, &newConfigs)
	if err != nil {
		return nil, err
	}
	return &newConfigs, nil
}
