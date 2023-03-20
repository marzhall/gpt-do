package main
import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"
        "io"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func getUserHomeDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
	    fmt.Println("Cannot get the current user to grab their home directory path:", err)
	    return ""
	}
    
	return currentUser.HomeDir
}

func getEnvDirectories() []string {
	homeEnv := fmt.Sprintf("%s/.env", getUserHomeDirectory())
	cwdEnv := ".env"
	return []string{homeEnv,cwdEnv}
}

func main() {
	godotenv.Load(getEnvDirectories()...)

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

        formattedMessage := fmt.Sprintf("%s for the following text, please and thank you: %s", prompt, data)
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