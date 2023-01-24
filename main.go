package main

import (
	"context"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("API KEY Missing")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	constant input File = "./input.txt"
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("failed to read file: %v", err)
	}
	msgPrefix := "give me a short list of libraries that are used in the code \n```python\n"
	msgSuffix := "\n```"
	msg := msgPrefix + string(fileBytes) + msgSuffix

	outputBuilder := strings.Builder{}
	// client.CompletionStreamWithEngine(ctx, gpt)

}
