package unusedparam

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/k0kubun/pp"
)

// Issue represents an issue found by linters.
type Issue struct {
	Text       string
	Pos        token.Position
	Suggestion *Suggestion
}

// Suggestion represents how to fix the issue.
type Suggestion struct {
	// Whether it needs to just delete without replacement codes.
	JustDelete bool
	NewLine    string
}

func Check(path string) ([]*Issue, error) {
	issues := []*Issue{}
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, path, src, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if dec, ok := n.(*ast.FuncDecl); ok {
			paramsMap := map[string]interface{}{}
			// TODO: Check if nil.
			// Make a map whose keys are paramsMap of the function.
			for _, l := range dec.Type.Params.List {
				for _, name := range l.Names {
					paramsMap[name.String()] = nil
				}
			}

			// Check if the params are used by the function.
			for _, stmt := range dec.Body.List {
				pp.Println(stmt)
				switch stmt := stmt.(type) {
				case *ast.AssignStmt:
					for _, lh := range stmt.Lhs {
						switch lh := lh.(type) {
						case *ast.Ident:
							if _, ok := paramsMap[lh.String()]; ok {
								paramsMap[lh.String()] = struct{}{}
							}
						}
					}
					for _, rh := range stmt.Rhs {
						switch rh := rh.(type) {
						case *ast.Ident:
							if _, ok := paramsMap[rh.String()]; ok {
								paramsMap[rh.String()] = struct{}{}
							}
						case *ast.BinaryExpr:
							switch x := rh.X.(type) {
							case *ast.Ident:
								if _, ok := paramsMap[x.String()]; ok {
									paramsMap[x.String()] = struct{}{}
								}
							}
							switch y := rh.Y.(type) {
							case *ast.Ident:
								if _, ok := paramsMap[y.String()]; ok {
									paramsMap[y.String()] = struct{}{}
								}
							}
						}
					}
				case *ast.ReturnStmt:
					// TODO: Add cases.

				}
			}
			for k, v := range paramsMap {
				if v != nil {
					continue
				}
				// TODO: Convert issues, means fetching the token.Position.
				fmt.Printf("%s is unused\n", k)
			}
		}
		return true
	})
	return issues, nil
}
