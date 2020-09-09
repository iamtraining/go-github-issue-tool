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

func OpenEditor(file string) error {
	cmd := exec.Command("vim", file)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()

}

func TempFile(curr string) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return "", fmt.Errorf("failure while creating temporary file")
	}

	defer os.Remove(file.Name())

	if curr != "" {
		_, err = file.WriteString(curr)
	}

	if err = file.Close(); err != nil {
		return "", fmt.Errorf("cant close tmp file")
	}

	if err = OpenEditor(file.Name()); err != nil {
		return "", err
	}

	body, err := Parse(file.Name())
	if err != nil {
		return "", fmt.Errorf("cant read temp file: %w", err)
	}

	return body, nil
}

func CreateIssue() (*entity.Issue, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Title: ")
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

func Parse(file string) (string, error) {
	tmp, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer tmp.Close()

	lines := []string{}

	scanner := bufio.NewScanner(tmp)
	for scanner.Scan() {
		t := scanner.Text()

		if len(t) == 0 {
			lines = append(lines, "\n")
		} else if t[0] != '#' {
			lines = append(lines, t)
		}
	}

	return strings.Join(lines, "/n"), nil
}
