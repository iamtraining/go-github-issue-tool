package editor

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/iamtraining/go-github-issue-tool/entity"
)

func OpenEditor() {
	cmd := exec.Command("cmd", "tmp.txt")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Errorf("cannot run editor %w", err)
	}

}

func TempFile(curr string) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		fmt.Errorf("failure while creating temporary file")
	}

	defer os.Remove(file.Name())

	if curr != "" {
		_, err := file.WriteString(curr)
	}

	if err = file.Close(); err != nil {
		fmt.Errorf("cant close tmp file")
	}
}

func CreateIssue() (*entity.Issue, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Title: \n")
	scanner.Scan()
	title := scanner.Text()

	fmt.Printf("Body: ")
	body, err := TempFile("")
	if err != nil {
		return nil, err
	}
	fmt.Println(body)

	issue := entity.Issue{
		Title: title,
		Body:  body,
	}

	return &issue, nil
}

func EditIssue(curr *entity.Issue) (*entity.Issue, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("current title: %s\n", curr.Title)
	fmt.Println("new title: ")
	title := ""

	if scanner.Scan() {
		title = scanner.Text()
	}

	fmt.Printf("new issue body: ")
	body, err := TempFile(curr.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(body)

	issue := entity.Issue{
		Title: title,
		Body:  body,
	}

	return &issue, nil
}

func Parse(file string) {
	tmp, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer tmp.Close()

	lines := []string{}

	scanner := bufio.NewScanner(tmp)
	for scanner.Scan() {
		t := scanner.Text()
		if t[0] != '#' {
			lines = append(lines, t)
		}
		if len(t) == 0 {
			lines = append(lines, "\n")
		}
	}

	line := strings.Join(lines[:len(lines)-1], "\n")
	fmt.Println(line)
}
