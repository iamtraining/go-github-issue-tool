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
	fmt.Println(`this tool is needed to interract with github issues
	available commands to use the tool:
	[create] to create issue; [help create] for more details
	[update] to update issue; [help update] for more details
	[updatestate] to update the state; [help updatestate] for more details
	[read] to read the selected issue; [help read] for more details
	[delete] to delete the selected issue; [help delete] for more details
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
func deleteusage() {
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
		requests.Create(oauth, os.Args[2], os.Args[3])
	case "update":
		requests.Update(oauth, os.Args[2], os.Args[3], os.Args[4])
	case "updatestate":
		requests.UpdateState(oauth, os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	case "read":
		requests.Read(oauth, os.Args[2], os.Args[3], os.Args[4])
	case "delete":
		requests.Delete(oauth, os.Args[2], os.Args[3], os.Args[4])
	case "help":
		switch os.Args[2] {
		case "create":
			createusage()
			os.Exit(1)
		case "update":
			updateusage()
			os.Exit(1)
		case "updatestate":
			updatestateusage()
			os.Exit(1)
		case "read":
			readusage()
			os.Exit(1)
		case "delete":
			deleteusage()
			os.Exit(1)
		}
	}
	//requests.Update(oauth, "iamtraining", "go-github-issue-tool", "1")
}
