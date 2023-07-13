package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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

func (c Client) GetGroup(groupIdentifier string) (Group, error) {
	resource := fmt.Sprintf("groups/%s", groupIdentifier)
	u := fmt.Sprintf("%s/%s", c.baseUrl, resource)

	// Prepare the request
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return Group{}, err
	}

	// Add authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return Group{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Group{}, err
	}

	var group Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		return Group{}, err
	}

	return group, nil
}
