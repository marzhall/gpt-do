package main
import (
	"context"
	"fmt"
	"log"
	"os"
        "io"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

        if (len(os.Args) < 2) {
                log.Fatal("Need to pass a directive - e.g.: gpto 'tell me the number after 4'")
        }

        prompt := os.Args[1]

        stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	data := string(stdin)

        formattedMessage := fmt.Sprintf("%s for the following text: %s", prompt, data)
        msg := gpt3.ChatCompletionRequestMessage {
                Role: "user",
                Content: formattedMessage,
        }

	resp, err := client.ChatCompletion(ctx, 
                gpt3.ChatCompletionRequest {
                        Messages:    []gpt3.ChatCompletionRequestMessage{msg},
                        MaxTokens: len(formattedMessage)*2,
                        Stop:      []string{"."},
                })
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}