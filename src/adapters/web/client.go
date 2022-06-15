package web

import (
	"bytes"
	"encoding/json"
	"github.com/mikejeuga/ghi-cli/issues"
	"github.com/mikejeuga/ghi-cli/models"
	"log"
	"net/http"
	"time"
)

type GHClient struct {
	BaseURL string
	Config  *models.Config
	Caller  *http.Client
}

func NewGHClient(config *models.Config) *GHClient {
	baseURL := "https://api.github.com/repos/"
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   5 * time.Second,
	}
	return &GHClient{BaseURL: baseURL, Config: config, Caller: client}
}

func (c *GHClient) CreateIssue(issue issues.Issue) {
	url := c.BaseURL + c.Config.Username + c.Config.Repo

	jsonBody, err := json.Marshal(issue)
	if err != nil {
		panic("cannot Marshall the Issue")
	}

	reader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	req.Header.Set("Accept", c.Config.HAccept)
	req.Header.Set("Authorization", "token "+c.Config.Tk)

	_, err = c.Caller.Do(req)
	if err != nil {
		log.Fatal(err)
	}

}
