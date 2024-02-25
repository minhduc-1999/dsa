package simplifypath

import (
	"dsa/utils/stack"
	"strings"
)

// Title: Simplify Path

// Problem link: https://leetcode.com/problems/simplify-path

// Testcases:
// case 0: "/home/"
// case 1: "/../"
// case 2: "/home//foo/"

func simplifyPath(path string) string {
	pathStack := stack.NewStack[string](len(path))
	start := 0
	path = path + "/"
	for i, ch := range path {
		if ch == '/' {
			if i > start {
				dir := path[start:i]
				switch dir {
				case ".":
				case "..":
					if pathStack.Size() > 1 {
						pathStack.Pop()
					}
				default:
					pathStack.Push(dir)
				}
			}
			start = i + 1
		}
	}
	return "/" + strings.Join(pathStack.Data(), "/")
}
