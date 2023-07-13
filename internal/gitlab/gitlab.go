package gitlab

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseUrl string
	token   string
}

type Group struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Url      string    `json:"web_url"`
	Projects []Project `json:"projects"`
}

type Project struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Url           string `json:"web_url"`
	MergeMethod   string `json:"merge_method"`
	DefaultBranch string `json:"default_branch"`
	Visibility    string `json:"visibility"`
}

func NewClient(baseUrl string, token string) Client {
	return Client{baseUrl, token}
}

func (c Client) Auth() error {
	// Set resource url
	resource := "projects"
	u := fmt.Sprintf("%s/%s", c.baseUrl, resource)

	// Prepare the request
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	// Add authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// Add query parameters
	q := req.URL.Query()
	q.Add("per_page", "1")
	req.URL.RawQuery = q.Encode()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return errors.New("non-2XX http status code")
	}

	return nil
}
