package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	instanceID       string
	apiTokenInstance string
	apiURL           string
	http             *http.Client
}

func NewClient(instanceID, apiToken, baseURL string) *Client {
	return &Client{
		instanceID:       instanceID,
		apiTokenInstance: apiToken,
		apiURL:           baseURL,
		http:             &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) SendMessage(chatID, message string) (*SendMessageResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", c.apiURL, c.instanceID, c.apiTokenInstance)

	body, err := json.Marshal(map[string]string{
		"chatId":  chatID,
		"message": message,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &SendMessageResponse{
		StatusCode: resp.StatusCode,
	}

	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			if err != io.EOF {
				return nil, err
			}
		}
		res.Error = &apiErr
		return res, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) GetChatHistory(chatID string, count int) (*GetChatHistoryResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getChatHistory/%s", c.apiURL, c.instanceID, c.apiTokenInstance)

	body, err := json.Marshal(map[string]any{
		"chatId": chatID,
		"count":  count,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &GetChatHistoryResponse{
		StatusCode: resp.StatusCode,
		Messages:   []Message{},
	}

	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			if err != io.EOF {
				return nil, err
			}
		}
		res.Error = &apiErr
		return res, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(&res.Messages); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) GetStateInstance() (*GetStateInstanceResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", c.apiURL, c.instanceID, c.apiTokenInstance)

	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &GetStateInstanceResponse{
		StatusCode: resp.StatusCode,
	}

	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			if err != io.EOF {
				return nil, err
			}
		}
		res.Error = &apiErr
		return res, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
