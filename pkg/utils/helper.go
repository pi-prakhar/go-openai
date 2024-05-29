package utils

import (
	"github.com/pi-prakhar/go-openai/pkg/logger"
	"github.com/pi-prakhar/utils/loader"
)

func FetchOpenAIAPIKey() string {
	openaiKey, err := loader.GetValueFromEnv("OPENAI_API_KEY")
	if err != nil {
		logger.Log.Error("Failed to load OPENAI_API_KEY", err)
	}
	return openaiKey
}

func IsProduction() bool {
	isProduction, err := loader.GetValueFromConf("production")
	if err != nil {
		logger.Log.Error("Failed to load key production from config files", err)
		return false
	}
	if isProduction == "true" {
		return true
	}
	return false
}

func Loadenv() {
	if !IsProduction() {
		err := loader.LoadEnv()
		if err != nil {
			logger.Log.Error("Failed to load env", err)
		}
		logger.Log.Info("Succesfully loaded local .env")
	}
}
