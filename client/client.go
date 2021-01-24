package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	host   string
	port   int
	client *http.Client
}

func NewClient(host string, port int) *Client {
	return &Client{
		host:   host,
		port:   port,
		client: &http.Client{},
	}
}

func (c *Client) get() (*[]Todo, error) {
	body, err := c.httpRequest("GET", "todo", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = json.NewDecoder(body).Decode(&todos)

	if err != nil {
		return nil, err
	}

	return &todos, nil
}

func (c *Client) httpRequest(method string, path string, body bytes.Buffer) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.getAbsolutePath(path), &body)

	if err != nil {
		return nil, err
	}

	if method == "PUT" || method == "POST" {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		resBody := new(bytes.Buffer)
		_, err = resBody.ReadFrom(res.Body)

		if err != nil {
			return nil, fmt.Errorf("server responded with an error: %v", res.StatusCode)
		}

		return nil, fmt.Errorf("server responded with an error: %v - %s", res.StatusCode, resBody.String())
	}

	return res.Body, nil
}

func (c *Client) getAbsolutePath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.host, c.port, path)
}
