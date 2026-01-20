package llm

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/SaketV8/test-generator-ai/internal/utils"
	"github.com/sashabaranov/go-openai"

	_ "github.com/joho/godotenv/autoload"
)

func ChatWithLLM(messageRoleSystem, messageRoleUser string, writeToLog bool) (string, error) {

	token := os.Getenv("GITHUB_PAT")
	if token == "" {
		return "", errors.New("environment variable not set")
	}

	// Client Config
	// to support the Github Hosted models
	config := openai.DefaultConfig(token)
	config.BaseURL = "https://models.github.ai/inference"

	// client with the custom configuration
	client := openai.NewClientWithConfig(config)

	// chat completion request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "openai/gpt-4.1-mini",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an expert Go developer assistant that generates file content based on user specifications.",
					// Content: messageRoleSystem,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "What are your capabilities",
					// Content: prompts.DefaultPrompt,
					// Content: messageRoleUser,
				},
			},
			// Temperature: 1.0,
			Temperature: 0.1,
			// TopP:        1.0,
		},
	)

	// Check for errors
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	llmOutput := resp.Choices[0].Message.Content
	if writeToLog {
		err = utils.WriteFile("./llmOutput/generated.txt", llmOutput)
		if err != nil {
			log.Fatalf("error writing file: %v", err)
		}
	}

	return llmOutput, nil
}
