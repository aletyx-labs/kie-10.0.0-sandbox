// File: main.go
package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Metadata holds configuration values
type Metadata struct {
	Port string
}

// ValidationResult represents the result of commit message validation
type ValidationResult struct {
	Result  bool     `json:"result"`
	Reasons []string `json:"reasons,omitempty"`
}

// Validate checks if a commit message meets the required criteria
func Validate(message string) ValidationResult {
	messageLength := len(message)
	var reasons []string

	// Check if message is too short (less than 5 characters)
	if messageLength < 5 {
		reasons = append(reasons, "Commit message is too short. Minimum length is 5 characters.")
	}

	// Check if message is too long (more than 72 characters)
	if messageLength > 72 {
		reasons = append(reasons, "Commit message is too long. Maximum length is 72 characters.")
	}

	// Message is valid if there are no reasons for failure
	return ValidationResult{
		Result:  len(reasons) == 0,
		Reasons: reasons,
	}
}

func main() {
	// Set up metadata
	metadata := Metadata{
		Port: os.Getenv("PORT"),
	}

	// Default to port 8080 if not specified
	if metadata.Port == "" {
		metadata.Port = "3000"
	}

	// Set up Gin router
	router := gin.Default()
	router.Use(cors.Default())

	// Define validation endpoint
	router.POST("/validate", func(context *gin.Context) {
		body, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, "Wrong input")
			return
		}
		context.JSON(http.StatusOK, Validate(string(body)))
	})

	// Start the server
	router.Run(":" + metadata.Port)
}
