package ask

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"time"
)

type CaffeineIntake struct {
	Timestamp    time.Time `json:"timestamp"`
	CaffeineInMg int       `json:"caffeineInMg"`
}

func (cs CaffeineIntake) MarshalJSON() ([]byte, error) {
	schema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"timestamp": map[string]interface{}{
				"type":        "string",
				"format":      "date-time",
				"description": "The time of caffeine consumption",
			},
			"caffeineInMg": map[string]interface{}{
				"type":        "integer",
				"description": "The amount of caffeine consumed in milligrams",
			},
		},
		"required": []string{"timestamp", "caffeineInMg"},
	}
	return json.Marshal(schema)
}

type Client struct {
	client openai.Client
}

func New() *Client {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &Client{client: *client}
}

// Ask takes the string input and throws it against the OpenAI API in expectation
// of geting the caffaine intake in milligram and a timestamp.
func (c *Client) Ask(userInput string) (CaffeineIntake, error) {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		MaxTokens:   500,
		Temperature: 0.4,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: fmt.Sprintf(systemMessage, time.Now()),
			},
			{
				Role:    "user",
				Content: userInput,
			},
		},
	}
	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return CaffeineIntake{}, err
	}

	// Create an instance of your struct
	var answer CaffeineIntake

	// Unmarshal the JSON into the struct
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &answer)
	if err != nil {
		fmt.Println("Error:", err)
		return CaffeineIntake{}, err
	}
	return answer, nil
}

var systemMessage = "You are an assistant that provides information about caffeine consumption. " +
	"\nYour responses should always be in JSON format with the following structure:\n{\n  \"timestamp\": ,\n  \"caffeineInMg\": \n}\nWhere \"timestamp\" is the ISO 8601 formatted date and time of consumption, \nand \"caffeineInMg\" is an integer representing the amount of caffeine in milligrams. The time right now is: %v"
