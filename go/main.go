package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

type Query struct {
	Variables map[string]interface{} `json:"variables"`
	Query     string                 `json:"query"`
}

type RawProblem struct {
	Data Data `json:"data"`
}

const LEETCODE_URL = "https://leetcode.com"

type Data struct {
	Question Question `json:"question"`
}

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (p Param) ParseGoDataType() string {
	switch p.Type {
	// int
	case "integer":
		return "int"
	case "integer[]":
		return "[]int"
	case "integer[][]":
		return "[][]int"
		// string
	case "string":
		return "string"
	case "string[]":
		return "[]string"
	case "string[][]":
		return "[][]string"
		//char
	case "character":
		return "byte"
	case "character[]":
		return "[]byte"
	case "character[][]":
		return "[][]byte"
	default:
		return ""
	}
}

type Return struct {
	Type string `json:"type"`
}

type MetaData struct {
	Name   string  `json:"name"`
	Params []Param `json:"params"`
	Return Return  `json:"return"`
}

type CodeSnippet struct {
	Code     string `json:"code"`
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
}

type Question struct {
	CodeSnippets        []CodeSnippet `json:"codeSnippets"`
	ExampleTestcaseList []string      `json:"exampleTestcaseList"`
	MetaData            string        `json:"metaData"`
	FrontendId          string        `json:"questionFrontendId"`
	Title               string        `json:"questionTitle"`
	CanSee              bool          `json:"canSeeQuestion"`
	TitleSlug           string        `json:"questionTitleSlug"`
	IsPaidOnly          bool          `json:"isPaidOnly"`
}

func (q Question) CodeSnippet(lang string) string {
	for _, snippet := range q.CodeSnippets {
		if snippet.Lang == lang {
			return snippet.Code
		}
	}
	return ""
}

func (q Question) ParseMetaData() (MetaData, error) {
	var meta MetaData
	err := json.Unmarshal([]byte(q.MetaData), &meta)
	if err != nil {
		return MetaData{}, err
	}
	return meta, nil
}

func (q Question) TransformTitleSlug() string {
	return strings.ReplaceAll(q.TitleSlug, "-", "_")
}

func (q Question) ProblemLink() string {
	return fmt.Sprintf("%s/problem/%s", LEETCODE_URL, q.TitleSlug)
}

func (q Question) SolutionFileName() string {
  return path.Join(q.Dirname(), fmt.Sprintf("%s.go", q.TransformTitleSlug()))
}

func (q Question) TestFileName() string {
	return fmt.Sprintf("%s_test.go", q.TransformTitleSlug())
}

func (q Question) Dirname() string {
	return path.Join(SOLUTION_DIR, fmt.Sprintf("%s_%s", q.FrontendId, q.TransformTitleSlug()))
}

func (q Question) PackageName() string {
	return strings.ReplaceAll(strings.ToLower(q.TitleSlug), "-", "")
}

func (q Question) Testcases() []string {
  var result []string
  for _, testcase := range q.ExampleTestcaseList {
    result = append(result, strings.ReplaceAll(testcase, "\n", "\t"))
  }
  return result
}

func main() {
	// question, err := fetchProblem("reverse-words-in-a-string")
	question, err := fetchProblem("two-sum")
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
}

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
