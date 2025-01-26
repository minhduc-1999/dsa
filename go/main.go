package main

import (
	"dsa/fetcher"
	"dsa/generator"
	"fmt"
	"os"
)

const TEMPLATE_PATH = "./template.txt"
const SOLUTION_DIR = "solutions"

type AppArgs struct {
	QuestionSlug string
	PackageName  string
}

func parseArgs() AppArgs {
	args := os.Args
	if len(args) < 2 {
		panic("Required arguments: slug")
	}
	packageName := args[1]
	if len(args) > 2 {
		packageName = args[2]
	}
	return AppArgs{
		QuestionSlug: args[1],
		PackageName:  packageName,
	}
}

func main() {
	appArgs := parseArgs()
	question, err := fetcher.FetchProblem(appArgs.QuestionSlug)
	if err != nil {
		panic(err.Error())
	}
	if !question.CanSee || question.IsPaidOnly {
		panic(fmt.Sprintf("Cannot view question %v\n", question.TitleSlug))
	}
	g := generator.NewGenerator(appArgs.PackageName, TEMPLATE_PATH, SOLUTION_DIR, question)
	err = g.Generate()
	if err != nil {
		panic(err.Error())
	}
}
