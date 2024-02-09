package main

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
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
