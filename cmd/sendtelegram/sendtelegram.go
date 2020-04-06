package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// env variable holding the API token
const envAPIToken = "SENDTELEGRAM_API_TOKEN"

// env variable holding the chat ID token
const envChatID = "SENDTELEGRAM_CHAT_ID"

// hard coded endpoint url
const endpoint = "https://api.telegram.org/bot%s/sendMessage"

// message length
const messageLen = 4096

// send message to a Telegram chatID using the input API token.
func sendMessage(apiToken string, chatID string) {
	// set chat endpoint
	chatEndpoint := fmt.Sprintf(endpoint, apiToken)
	// reading from stdin for easy pipeing
	reader := bufio.NewReader(os.Stdin)
	// buffering reads. currently telegram max message length is 4096 bytes
	buffer := make([]byte, messageLen)

	for {
		// read input
		n, err := io.ReadFull(reader, buffer)
		// stop if finished reading or errors occur.
		if err != nil && err == io.EOF {
			break
		}
		// build HTTP request.
		resp, err := http.PostForm(chatEndpoint, url.Values{"chat_id": {chatID}, "text": {string(buffer[:n])}})
		// if response is not HTTP.OK, log error and exit.
		if resp.StatusCode != 200 {
			log.Fatal(fmt.Sprintf("Error while sending message. HTTP response: %d (error: %v)", resp.StatusCode, err))
		}
	}
}

func main() {
	apiToken := flag.String("api-token", os.Getenv(envAPIToken), "Telegram API token.")
	chatID := flag.String("chat-id", os.Getenv(envChatID), "Telegram chat ID.")
	flag.Parse()

	// checking if the token and chat id are present
	if (*apiToken == "") || (*chatID == "") {
		log.Fatal("API token or chat id are not set.")
	}
	// send message
	sendMessage(*apiToken, *chatID)
}
