package generator

import (
	. "dsa/fetcher"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

type generator struct {
	packageName  string
	templatePath string
	solutionDir  string
	question     Question
}

func NewGenerator(packageName string, templatePath string, solutionDir string, question Question) generator {
	return generator{
		packageName,
		templatePath,
		solutionDir,
		question,
	}
}

func (g generator) Generate() error {
	err := g.generateSolutionTemplate()
	if err != nil {
		return err
	}
	err = g.generateTestTemplate()
	if err != nil {
		return err
	}
	g.printUsefullCommands()
	return nil
}

func (g generator) printUsefullCommands() {
	fmt.Printf("Commands:\n\tRun test:\t%v\n", g.TestCmd())
}

func (g generator) generateTestTemplate() error {
	meta, _ := g.question.ParseMetaData()
	cmd := exec.Command("gotests", "-w", "-only", fmt.Sprintf("^%s$", meta.Name), g.SolutionFileName())
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (g generator) generateSolutionTemplate() error {
	if _, err := os.Stat(g.Dirname()); err == nil {
		return fmt.Errorf("problem %v has already solved", g.question.Title)
	}
	err := os.Mkdir(g.Dirname(), os.ModePerm)
	if err != nil {
		return err
	}
	solutionFile, err := os.Create(g.SolutionFileName())
	if err != nil {
		return err
	}
	defer solutionFile.Close()
	tmpl, err := template.ParseFiles(g.templatePath)
	if err != nil {
		return err
	}
	err = tmpl.Execute(solutionFile, TemplateOpts{
		PackageName: g.PackageName(),
		Title:       g.question.Title,
		ProblemLink: g.ProblemLink(),
		Testcases:   g.question.Testcases(),
		CodeSnippet: g.question.CodeSnippet("Go"),
	})
	return err
}

func (g generator) ProblemLink() string {
	return fmt.Sprintf("%s/problems/%s", LEETCODE_URL, g.question.TitleSlug)
}

func (g generator) SolutionFileName() string {
	return path.Join(g.Dirname(), fmt.Sprintf("%s.go", g.question.TransformTitleSlug()))
}

func (g generator) TestFileName() string {
	return fmt.Sprintf("%s_test.go", g.question.TransformTitleSlug())
}

func (g generator) Dirname() string {
	return path.Join(g.solutionDir, fmt.Sprintf("%s_%s", g.question.FrontendId, g.question.TransformTitleSlug()))
}

func (g generator) TestCmd() string {
	return fmt.Sprintf("go test %v", path.Join(DSA_MODULE, g.Dirname()))
}

func (g generator) PackageName() string {
	if len(g.packageName) > 0 {
		return g.packageName
	}
	return strings.ReplaceAll(strings.ToLower(g.question.TitleSlug), "-", "")
}
