package issues

type Issue struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees"`
}

func NewIssue(title string) *Issue {
	return &Issue{Title: title, Body: "", Assignees: nil}
}

func (i *Issue) WriteBody(content string) {
	i.Body = content
}

func (i *Issue) AddAssignees(person string) {
	i.Assignees = append(i.Assignees, person)
}
