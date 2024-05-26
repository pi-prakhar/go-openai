package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	database "github.com/pi-prakhar/go-openai/internal/db"
	"github.com/pi-prakhar/go-openai/internal/models"
	"github.com/pi-prakhar/go-openai/pkg/logger"
	"github.com/pi-prakhar/go-openai/pkg/utils"
)

var OPENAI_CHAT_COMPLETION_URL string = "https://api.openai.com/v1/chat/completions"
var MODEL_GPT_3_5_TURBO string = "gpt-3.5-turbo"
var DEFAULT_TEMPERATURE float64 = 0.7

// var messages []models.ChatMessage

func GetMessages() *[]models.ChatMessage {
	return &database.Messages
}

func SendMessage(message string) (*models.ChatResponse, error) {
	database.Messages = append(database.Messages, *createMessage("user", &message))
	body, err := createRequestBody(&MODEL_GPT_3_5_TURBO, &DEFAULT_TEMPERATURE, &database.Messages)

	if err != nil {
		logger.Log.Debug("Error : Failed SendMessage")
		return nil, err
	}

	res, err := sendRequest(body)
	if err != nil {
		logger.Log.Debug("Error : Failed SendMessage")
		return nil, err
	}

	data, err := parseResponse(res)
	if err != nil {
		logger.Log.Debug("Error : Failed SendMessage")
		return nil, err
	}

	updateMessages(data)
	return data, nil

}
func updateMessages(data *models.ChatResponse) {
	var lastMessage models.ChatMessage = data.Choices[len(data.Choices)-1].Message
	database.Messages = append(database.Messages, lastMessage)
}

func parseResponse(r *http.Response) (*models.ChatResponse, error) {
	var resData models.ChatResponse
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log.Debug("Error : Failed to read response body")
		return nil, err
	}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		logger.Log.Debug("Error : Failed to unmarshal response body")
		return nil, err
	}
	return &resData, nil

}
func sendRequest(requestBody *[]byte) (*http.Response, error) {
	openAIAPIKey := utils.FetchOpenAIAPIKey()
	req, err := http.NewRequest("POST", OPENAI_CHAT_COMPLETION_URL, bytes.NewBuffer(*requestBody))
	if err != nil {
		logger.Log.Debug("Error : Failed to create request")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAIAPIKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Debug("Error : Failed to send request")
		return nil, err
	}
	defer req.Body.Close()
	return res, nil
}
func createMessage(role string, content *string) *models.ChatMessage {
	message := models.ChatMessage{
		Role:    role,
		Content: *content,
	}
	return &message
}

func createRequestBody(model *string, temperature *float64, messages *[]models.ChatMessage) (*[]byte, error) {
	var requestBody = models.ChatRequest{
		Model:       *model,
		Messages:    *messages,
		Temperature: *temperature,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logger.Log.Debug("Error : Failed to marshal request body data to json")
		return nil, err
	}
	fmt.Printf(string(jsonData))
	byteData := []byte(jsonData)

	return &byteData, nil

}
