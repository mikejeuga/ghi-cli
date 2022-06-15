package issues

import (
	"log"
	"os"
)

type Issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees"`
}

func NewIssue(title string) *Issue {
	return &Issue{Title: title, Body: "", Assignees: nil}
}

func (i *Issue) WriteBody(content string) error {
	file, err := os.ReadFile("./" + content + ".md")
	if err != nil {
		log.Fatal(err)
	}
	i.Body = string(file)
	return nil
}

func (i *Issue) AddAssignees(person string) {
	i.Assignees = append(i.Assignees, person)
}
