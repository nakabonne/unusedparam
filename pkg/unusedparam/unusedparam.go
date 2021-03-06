package unusedparam

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

const blankIdentifier = "_"

// Issue represents an issue found by linters.
type Issue struct {
	Text string
	Pos  token.Position
}

func (i *Issue) String() string {
	return fmt.Sprintf("%s:%d:%d: %s", i.Pos.Filename, i.Pos.Line, i.Pos.Column, i.Text)
}

// parameter represents a function parameter.
type parameter struct {
	ident *ast.Ident
	used  bool
}

// Check inspects a single file.
func Check(path string) ([]*Issue, error) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, path, src, parser.ParseComments)
	if len(f.Comments) > 0 && isGenerated(f.Comments[0].Text()) {
		return nil, fmt.Errorf("%s is a generated file", path)
	}

	issues := []*Issue{}
	ast.Inspect(f, func(n ast.Node) bool {
		dec, ok := n.(*ast.FuncDecl)
		if !ok || dec.Body == nil {
			return true
		}

		// Make a map whose keys are names of params.
		paramsMap := map[string]*parameter{}
		for _, l := range dec.Type.Params.List {
			for _, n := range l.Names {
				paramsMap[n.String()] = &parameter{ident: n}
			}
		}

		// Check if the params are used by any of statement nodes.
		for _, stmt := range dec.Body.List {
			ast.Inspect(stmt, func(n ast.Node) bool {
				switch n := n.(type) {
				case *ast.Ident:
					if _, ok := paramsMap[n.String()]; ok {
						paramsMap[n.String()].used = true
					}
				}
				return true
			})
		}

		// Make a slice of issues based on the unused params.
		for name, param := range paramsMap {
			if param.used {
				continue
			}
			if name == blankIdentifier {
				continue
			}
			issues = append(issues, &Issue{
				Text: fmt.Sprintf("%s is unused in %s", name, dec.Name.String()),
				Pos:  fset.Position(param.ident.Pos()),
			})
		}
		return true
	})
	return issues, nil
}

// isGenerated checks if a given file is generated by using a comment text.
func isGenerated(text string) bool {
	return strings.Contains(text, "Code generated") || strings.Contains(text, "DO NOT EDIT")
}
