package main

import (
	"fmt"
	"os"

	"github.com/iamtraining/go-github-issue-tool/requests"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("couldnt load .env")
	}
}

func usage() {
	fmt.Println(`
	this tool is needed to interract with github issues
	available commands to use the tool:
	[create] to create issue
	[update] to update issue
	[updatestate] to update the state
	[read] to read the selected issue
	`)
}

func createusage() {
	fmt.Println(`
	enter 2 arguments: <user> <repo>`)
}
func updateusage() {
	fmt.Println(`enter 3 arguments: <user> <repo> <number>`)
}
func updatestateusage() {
	fmt.Println(`enter 4 arguments: <user> <repo> <number> <state>
	state should only have 2 conditions <open> or <closed>`)
}
func readusage() {
	fmt.Println(`enter 3 arguments: <user> <repo> <number>`)
}

func main() {
	oauth := os.Getenv("OAUTH")
	if oauth == "" {
		fmt.Println("you havent specified oauth token in .env file")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "create":
		if len(os.Args) < 4 {
			createusage()
			os.Exit(1)
		}
		requests.Create(oauth, os.Args[2], os.Args[3])
	case "update":
		if len(os.Args) < 5 {
			updateusage()
			os.Exit(1)
		}
		requests.Update(oauth, os.Args[2], os.Args[3], os.Args[4])
	case "updatestate":
		if len(os.Args) < 6 {
			updatestateusage()
			os.Exit(1)
		}
		requests.UpdateState(oauth, os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	case "read":
		if len(os.Args) < 5 {
			readusage()
			os.Exit(1)
		}
		requests.Read(oauth, os.Args[2], os.Args[3], os.Args[4])
	case "help":
		usage()
		os.Exit(1)
	}
	//requests.Update(oauth, "iamtraining", "go-github-issue-tool", "1")
}
