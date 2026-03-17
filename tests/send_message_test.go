package tests

import (
	"green-api-test/client"
	"green-api-test/config"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	cfg = config.Init()
	os.Exit(m.Run())
}

func TestSendMessage_Success(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)

	resp, err := sendMessageWithRetry(t, client, cfg.ChatID, "test")
	require.NoError(t, err)

	require.Equal(t, 200, resp.StatusCode)
	require.NotEmpty(t, resp.MessageID)
}

func TestSendMessage_EmptyChatID(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)
	resp, err := sendMessageWithRetry(t, client, "", "test")
	require.NoError(t, err)

	require.Equal(t, 400, resp.StatusCode)
}

func TestSendMessage_EmptyMessage(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)
	resp, err := sendMessageWithRetry(t, client, cfg.ChatID, "")
	require.NoError(t, err)

	require.Equal(t, 400, resp.StatusCode)
}

func sendMessageWithRetry(t *testing.T, c *client.Client, chatID string, message string) (*client.SendMessageResponse, error) {
	t.Helper()

	var resp *client.SendMessageResponse
	var err error

	for i := 0; i < 3; i++ {
		resp, err = c.SendMessage(chatID, message)
		require.NoError(t, err)

		if resp.StatusCode != 429 {
			return resp, err
		}

		time.Sleep(1 * time.Second)
	}

	t.Fatalf("failed due to rate limit (429)")
	return nil, nil
}
