package issues_test

import (
	"github.com/mikejeuga/ghi-cli/issues"
	"github.com/mikejeuga/ghi-cli/testhelpers"
	"testing"
)

func TestIssue(t *testing.T) {
	t.Run("I can write a body of the md file to the issue", func(t *testing.T) {
		issue := issues.NewIssue("Test Title")
		body := `# What ?
We want to make a GH issue for code reviews

# How ?
- [ ] We will use a cli application
- [ ] Give the title
- [ ] enter the name of the md for the body
- [ ] Assign the issue to someone, yourself by default
- [ ] Call the http client

# Why ?
## It is fun`

		err := issue.WriteBody("init")
		testhelpers.AssertError(t, err)
		testhelpers.AssertEqual(t, issue.Body, body)
	})

	t.Run("Can add an assignee", func(t *testing.T) {
		issue := issues.NewIssue("Test Title for Assignee")
		assignee := "Mike la menace"

		issue.AddAssignees(assignee)

		testhelpers.AssertEqual(t, len(issue.Assignees), 1)
		testhelpers.AssertEqual(t, issue.Assignees[0], assignee)
	})

}
