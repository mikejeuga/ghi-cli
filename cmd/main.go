package main

import (
	"github.com/mikejeuga/ghi-cli/issues"
	"github.com/mikejeuga/ghi-cli/models"
	"github.com/mikejeuga/ghi-cli/src/adapters/web"
	"log"
)

func main() {
	config := models.NewConfig()

	issue := issues.NewIssue("Test 4th time")

	err := issue.WriteBody("init")
	if err != nil {
		log.Fatal(err)
	}
	issue.AddAssignees("")

	ghClient := web.NewGHClient(config)
	ghClient.CreateIssue(*issue)

}
