package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamtraining/go-github-issue-tool/editor"
	"github.com/iamtraining/go-github-issue-tool/entity"
)

const (
	repoUser = "https://api.github.com/repos/%s/%s/issues"
	issueNum = repoUser + "/%s"
)

func sendRequest(oauth, method, url string, issue *entity.Issue) (*entity.Issue, error) {
	marshal, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("failure while: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(marshal))
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", oauth))
	if err != nil {
		return nil, fmt.Errorf("creating request failure: %w", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending request %w", err)
	}

	defer resp.Body.Close()

	result := entity.Issue{}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode failure: %w", err)
	}

	return &result, nil
}

func Create(oauth, user, repo string) {
	issue, err := editor.CreateIssue()
	if err != nil {
		fmt.Println(err)
	}
	_, err = sendRequest(oauth, "POST", fmt.Sprintf(repoUser, user, repo), issue)
	if err != nil {
		fmt.Println(err)
	}
}

func Read(oauth, user, repo, number string) {
	result, err := sendRequest(oauth, "GET", fmt.Sprintf(issueNum, user, repo, number), nil)
	if err != nil {
		fmt.Println(err)
	}
	result.String()
}

func Update(oauth, user, repo, number string) {
	result, err := sendRequest(oauth, "GET", fmt.Sprintf(issueNum, user, repo, number), nil)
	if err != nil {
		fmt.Println(err)
	}
	issue, err := editor.EditIssue(result)
	result, err = sendRequest(oauth, "PATCH", fmt.Sprintf(issueNum, user, repo, number), issue)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateState(oauth, user, repo, number, state string) {
	switch state {
	case "closed", "open":
		issue := entity.Issue{
			State: state,
		}
		_, err := sendRequest(oauth, "PATCH", fmt.Sprintf(issueNum, user, repo, number), &issue)
		if err != nil {
			fmt.Println(err)
		}
	default:
		return
	}

}

func Delete(oauth, user, repo, number string) {
	_, err := sendRequest(oauth, "DELETE", fmt.Sprintf(issueNum, user, repo, number), nil)
	if err != nil {
		fmt.Println(err)
	}
}
