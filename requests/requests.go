package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/iamtraining/go-github-issue-tool/entity"
)

func makeRequest(method, url, token string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant send HTTP request" + err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode < 300 {
		return nil, fmt.Errorf("soemthing goes wrong")
	}

	parse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant parse resp body")
	}

	return parse, nil
}

func CreateIssue() {
	oauth := os.Getenv("OAUTH")

	issue := entity.Issue{
		Title: "testing1212",
		Body:  "rawr xd",
	}
	marshal, err := json.Marshal(issue)
	if err != nil {
		panic("cannot marshal")
	}

	req, err := http.NewRequest("POST", "https://api.github.com/repos/iamtraining/go-github-issue-tool/issues", bytes.NewBuffer(marshal))
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauth))
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failure while getting response %w", err)
	}

	defer resp.Body.Close()

	m := map[string]interface{}{}

	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		fmt.Errorf("decode failure %w", err)
	}
	fmt.Println(m)
}

func OpenIssue() {
	oauth := os.Getenv("OAUTH")

	issue := entity.Issue{
		State: "open",
	}
	marshal, err := json.Marshal(issue)
	if err != nil {
		panic("cannot marshal")
	}

	req, err := http.NewRequest("PATCH", "https://api.github.com/repos/iamtraining/go-github-issue-tool/issues/1", bytes.NewBuffer(marshal))
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauth))
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failure while getting response %w", err)
	}

	defer resp.Body.Close()

	m := map[string]interface{}{}

	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		fmt.Errorf("decode failure %w", err)
	}
	fmt.Println(m)
}

func CloseIssue() {
	oauth := os.Getenv("OAUTH")

	issue := entity.Issue{
		State: "closed",
	}
	marshal, err := json.Marshal(issue)
	if err != nil {
		panic("cannot marshal")
	}

	req, err := http.NewRequest("PATCH", "https://api.github.com/repos/iamtraining/go-github-issue-tool/issues/1", bytes.NewBuffer(marshal))
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauth))
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failure while getting response %w", err)
	}

	defer resp.Body.Close()

	m := map[string]interface{}{}

	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		fmt.Errorf("decode failure %w", err)
	}
	fmt.Println(m)
}

func EditIssue() {
	oauth := os.Getenv("OAUTH")

	issue := entity.Issue{
		Title: "editing title",
		Body:  "editing body",
	}
	marshal, err := json.Marshal(issue)
	if err != nil {
		panic("cannot marshal")
	}

	req, err := http.NewRequest("PATCH", "https://api.github.com/repos/iamtraining/go-github-issue-tool/issues/1", bytes.NewBuffer(marshal))
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauth))
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("failure while getting response %w", err)
	}

	defer resp.Body.Close()

	m := map[string]interface{}{}

	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		fmt.Errorf("decode failure %w", err)
	}
	fmt.Println(m)
}
