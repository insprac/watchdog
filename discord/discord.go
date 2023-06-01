package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/insprac/watchdog/config"
	"github.com/insprac/watchdog/utils"
)

type DiscordPayload struct {
	Content string `json:"content"`
}

func SendWebhook(message string) error {
	payload := DiscordPayload{
		Content: utils.MarkNumbersBoldMarkdown(message),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling Discord payload: %v", err)
	}

	resp, err := http.Post(config.GetDiscordWebhookUrl(), "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending Discord notification: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("received non-2xx response code: %d", resp.StatusCode)
	}

	return nil
}
