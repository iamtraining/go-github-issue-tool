package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iamtraining/go-github-issue-tool/entity"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("couldnt load .env")
	}
}

func main() {
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
