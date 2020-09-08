package main

import (
	"os"

	"github.com/iamtraining/go-github-issue-tool/requests"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("couldnt load .env")
	}
}

func main() {
	oauth := os.Getenv("OAUTH")
	requests.Create(oauth, "iamtraining", "go-github-issue-tool")
}
