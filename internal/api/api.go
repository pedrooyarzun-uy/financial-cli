package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var CLIENT *ApiClient

type ApiClient struct {
	BaseURL    string
	HttpClient *http.Client
	Token      string
}

func Init(baseUrl string) {
	CLIENT = &ApiClient{
		BaseURL: baseUrl,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *ApiClient) SetToken(token string) {
	c.Token = token
}

func (c *ApiClient) PostMethod(endpoint string, result interface{}, payload interface{}, public bool) error {
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	if !public {
		req.Header.Set("Authorization", c.Token)
	}

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (c *ApiClient) GetMethod(endpoint string, result interface{}) error {

	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Token)

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)

}
