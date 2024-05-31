package main

import (
	"net/http"
	"time"

	"github.com/pi-prakhar/go-openai/api"
	"github.com/pi-prakhar/go-openai/pkg/logger"
	"github.com/pi-prakhar/go-openai/pkg/utils"
	"github.com/rs/cors"
)

func init() {
	logger.InitLogger()
	logger.Log.Info("Logger started")
	utils.Loadenv()
}

func main() {
	//cors settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000/", "http://localhost:8000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})
	handler := c.Handler(api.Router())

	svr := &http.Server{
		Addr:           ":8000",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Log.Error("Failed to start server", svr.ListenAndServe())
}
