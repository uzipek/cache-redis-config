package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadRedisConfig(file string) (*map[string]interface{}, error) {
	var config map[string]interface{}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", file)
	}
	data, err := os.ReadFile(file)
	if err!= nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &config); err!= nil {
		return nil, err
	}
	return &config, nil
}

func GetRedisPassword(config *map[string]interface{}) string {
	if redisPassword, ok := (*config)["redis_password"].(string); ok {
		return redisPassword
	}
	return ""
}

func GetRedisHost(config *map[string]interface{}) string {
	if redisHost, ok := (*config)["redis_host"].(string); ok {
		return redisHost
	}
	return ""
}

func GetRedisPort(config *map[string]interface{}) int {
	if redisPort, ok := (*config)["redis_port"].(int); ok {
		return redisPort
	}
	return 6379
}

func GetRedisDB(config *map[string]interface{}) int {
	if redisDB, ok := (*config)["redis_db"].(int); ok {
		return redisDB
	}
	return 0
}

func GetRedisTimeout(config *map[string]interface{}) int {
	if redisTimeout, ok := (*config)["redis_timeout"].(int); ok {
		return redisTimeout
	}
	return 0
}

func GetRedisPrefix(config *map[string]interface{}) string {
	if redisPrefix, ok := (*config)["redis_prefix"].(string); ok {
		return redisPrefix
	}
	return ""
}

func GetRedisUsername(config *map[string]interface{}) string {
	if redisUsername, ok := (*config)["redis_username"].(string); ok {
		return redisUsername
	}
	return ""
}

func GetRedisPasswordFile(config *map[string]interface{}) string {
	if redisPasswordFile, ok := (*config)["redis_password_file"].(string); ok {
		return redisPasswordFile
	}
	return ""
}