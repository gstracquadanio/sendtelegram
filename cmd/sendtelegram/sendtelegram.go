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

// hard coded endpoint url
const endpoint = "https://api.telegram.org/bot%s/sendMessage"

// send message to the telegram chat id
func sendMessage(apiToken string, chatID string) {
	// set chat endpoint
	chatEndpoint := fmt.Sprintf(endpoint, apiToken)
	// reading from stdin for easy pipeing
	reader := bufio.NewReader(os.Stdin)
	// buffering reads. currently telegram max message length is 4096 bytes
	buffer := make([]byte, 4096)

	for {
		// read input
		_, err := reader.Read(buffer)
		// stop if finished reading or errors occur.
		if err != nil && err == io.EOF {
			break
		}
		// build HTTP request.
		resp, err := http.PostForm(chatEndpoint, url.Values{"chat_id": {chatID}, "text": {string(buffer)}})
		// if response is not HTTP.OK, log error and exit.
		if resp.StatusCode != 200 {
			log.Fatal(fmt.Sprintf("Error while sending message. HTTP response: %d (error: %v)", resp.StatusCode, err))
		}
	}
}

func main() {
	apiToken := flag.String("api-token", os.Getenv("SENDTELEGRAM_API_TOKEN"), "Telegram API token.")
	chatID := flag.String("chat-id", os.Getenv("SENDTELEGRAM_CHAT_ID"), "Telegram chat ID.")
	flag.Parse()

	// checking if the token and chat id are present
	if (*apiToken == "") || (*chatID == "") {
		log.Fatal("API token or chat id are not set.")
	}
	// send message
	sendMessage(*apiToken, *chatID)
}
