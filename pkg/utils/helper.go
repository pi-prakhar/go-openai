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
