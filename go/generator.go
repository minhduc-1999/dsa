package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

const TEMPLATE_PATH = "./template.txt"
const TEST_TEMPLATE_PATH = "./test_template.txt"
const SOLUTION_DIR = "solutions"

func generateTestTemplate(problem Question) error {
	meta, _ := problem.ParseMetaData()
	cmd := exec.Command("gotests", "-w", "-only", fmt.Sprintf("^%s$", meta.Name), problem.SolutionFileName())
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func generateSolutionTemplate(problem Question) error {
	if _, err := os.Stat(problem.Dirname()); err == nil {
		return fmt.Errorf("Problem %v has already solved", problem.Title)
	}
	err := os.Mkdir(problem.Dirname(), os.ModePerm)
	if err != nil {
		return err
	}
	solutionFile, err := os.Create(problem.SolutionFileName())
	defer solutionFile.Close()
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(TEMPLATE_PATH)
	if err != nil {
		return err
	}
	tmpl.Execute(solutionFile, problem)
	return nil
}

func fetchProblem(slug string) (Question, error) {
	const PROBLEM_URL = "https://leetcode.com/graphql/"
	const QUERY = `
  query questionData($titleSlug: String!) {
    question(titleSlug: $titleSlug) {
        questionFrontendId
        questionTitle
        questionTitleSlug
        metaData
        canSeeQuestion
        isPaidOnly
        exampleTestcaseList
        codeSnippets {
          lang
          langSlug
          code
        }
    }
  }
  `
	query_param := Query{
		Query: QUERY,
		Variables: map[string]interface{}{
			"titleSlug": slug,
		},
	}
	request_buf, err := json.Marshal(query_param)
	if err != nil {
		return Question{}, err
	}
	buf_reader := bytes.NewReader(request_buf)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, PROBLEM_URL, buf_reader)
	if err != nil {
		return Question{}, err
	}
	request.Header.Add("content-type", "application/json")
	request.Header.Add("referer", "https://leetcode.com")
	request.Header.Add("cookie", "csrftoken=Q70rHaUJFYC92Ey0ty3NuIRIkZCz15MCnMXUUsLj1FZwTnUVcC2GLOQ6bO2q37OT;")
	request.Header.Add("X-Csrftoken", "Q70rHaUJFYC92Ey0ty3NuIRIkZCz15MCnMXUUsLj1FZwTnUVcC2GLOQ6bO2q37OT")
	resp, err := client.Do(request)
	if err != nil {
		return Question{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Question{}, err
	}
	var rawProblem RawProblem
	err = json.Unmarshal(body, &rawProblem)
	if err != nil {
		return Question{}, err
	}
	return rawProblem.Data.Question, nil
}
