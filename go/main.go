package main

import (
	"fmt"
	"os"
)

type AppArgs struct {
	QuestionSlug string
}

func parseArgs() AppArgs {
	args := os.Args
	if len(args) < 2 {
		panic("Required question slug")
	}
	return AppArgs{
		QuestionSlug: args[1],
	}
}

func main() {
	appArgs := parseArgs()
	question, err := fetchProblem(appArgs.QuestionSlug)
	if err != nil {
		panic(err.Error())
	}
	if !question.CanSee || question.IsPaidOnly {
		panic(fmt.Sprintf("Cannot view question %v\n", question.TitleSlug))
	}
	err = generateSolutionTemplate(question)
	if err != nil {
		panic(err.Error())
	}
	err = generateTestTemplate(question)
	if err != nil {
		panic(err.Error())
	}
  printUsefullCommands(question)
}
