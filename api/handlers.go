package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pi-prakhar/go-openai/internal/models"
	"github.com/pi-prakhar/go-openai/pkg/logger"
	"github.com/pi-prakhar/go-openai/pkg/utils"
)

func HandleOpenAITest(w http.ResponseWriter, r *http.Request) {
	openAIAPIKey := utils.FetchOpenAIAPIKey()
	var response models.Responder
	reqBody := []byte(
		` {
				"model" : "gpt-3.5-turbo",
				"messages" : [{"role":"user", "content":"say this is as test!"}],
				"temperature" : 0.7
		}`,
	)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Log.Debug("Error : Failed to create request")
		response = models.ErrorResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
		response.WriteJSON(w, http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAIAPIKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Debug("Error : Failed to send request")
		response = models.ErrorResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
		response.WriteJSON(w, http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Log.Debug("Error : Failed to read response body")
		response = models.ErrorResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
		response.WriteJSON(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)

	_, err = w.Write(resBody)
	if err != nil {
		logger.Log.Debug("Error : Failed to write response")
		return
	}
}
