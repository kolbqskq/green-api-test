package tests

import (
	"green-api-test/client"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetChatHistory_Success(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)

	resp, err := getChatHistoryWithRetry(t, client, cfg.ChatID, 10)
	require.NoError(t, err)

	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, resp.Messages)

	require.Nil(t, resp.Error)

	for _, v := range resp.Messages {
		require.NotEmpty(t, v.MessageID)
		require.NotEmpty(t, v.ChatID)
		require.NotEmpty(t, v.Timestamp)
	}
}

func TestGetChatHistory_EmptyChatID(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)

	resp, err := getChatHistoryWithRetry(t, client, "", 10)
	require.NoError(t, err)

	require.NotNil(t, resp.Error)
	require.NotEmpty(t, resp.Error.Message)

	require.Equal(t, 400, resp.StatusCode)
	require.Contains(t, resp.Error.Message, "chatId")
}

func TestGetChatHistory_InvalidCount(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)

	resp, err := getChatHistoryWithRetry(t, client, cfg.ChatID, -1)
	require.NoError(t, err)

	require.NotNil(t, resp.Error)
	require.NotEmpty(t, resp.Error.Message)

	require.Equal(t, 400, resp.StatusCode)
	require.NotEmpty(t, resp.Error.Message)
	require.Contains(t, resp.Error.Message, "Count")
}

func TestGetChatHistory_LargeCount(t *testing.T) {
	client := client.NewClient(cfg.InstanceID, cfg.APITokenInstance, cfg.APIURL)

	requireAuthorized(t, client)

	resp, err := getChatHistoryWithRetry(t, client, cfg.ChatID, 1000000000000000000)
	require.NoError(t, err)

	require.NotNil(t, resp.Error)
	require.NotEmpty(t, resp.Error.Message)

	require.Equal(t, 400, resp.StatusCode)
	require.NotEmpty(t, resp.Error.Message)
	require.Contains(t, resp.Error.Message, "Count")
}

func getChatHistoryWithRetry(t *testing.T, c *client.Client, chatID string, count int) (*client.GetChatHistoryResponse, error) {
	t.Helper()

	var resp *client.GetChatHistoryResponse
	var err error

	for i := 0; i < 3; i++ {
		resp, err = c.GetChatHistory(chatID, count)
		require.NoError(t, err)

		if resp.StatusCode != 429 {
			return resp, err
		}

		time.Sleep(1 * time.Second)
	}

	t.Fatalf("failed due to rate limit (429)")
	return nil, nil
}
