package main

import (
	"github.com/mikejeuga/ghi-cli/issues"
	"github.com/mikejeuga/ghi-cli/models"
	"github.com/mikejeuga/ghi-cli/src/adapters/web"
)

func main() {
	config := models.NewConfig()

	issue := issues.NewIssue("Test 4th time")

	issue.WriteBody("# This is the 4th try, this one via go client")
	issue.AddAssignees("mikejeuga")

	ghClient := web.NewGHClient(config)
	ghClient.CreateIssue(*issue)

}
