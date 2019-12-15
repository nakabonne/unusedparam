package unusedparam

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

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

func Check(path string) ([]*Issue, error) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, path, src, parser.Mode(0))

	issues := []*Issue{}
	ast.Inspect(f, func(n ast.Node) bool {
		if dec, ok := n.(*ast.FuncDecl); ok {
			paramsMap := map[string]*parameter{}
			// TODO: Check if nil.
			// Make a map whose keys are names of params.
			for _, l := range dec.Type.Params.List {
				for _, n := range l.Names {
					paramsMap[n.String()] = &parameter{ident: n}
				}
			}

			// Check if the params are used by the function.
			for _, stmt := range dec.Body.List {
				switch stmt := stmt.(type) {
				case *ast.AssignStmt:
					for _, lh := range stmt.Lhs {
						switch lh := lh.(type) {
						case *ast.Ident:
							if _, ok := paramsMap[lh.String()]; ok {
								paramsMap[lh.String()].used = true
							}
						}
					}
					for _, rh := range stmt.Rhs {
						switch rh := rh.(type) {
						case *ast.Ident:
							if _, ok := paramsMap[rh.String()]; ok {
								paramsMap[rh.String()].used = true
							}
						case *ast.BinaryExpr:
							switch x := rh.X.(type) {
							case *ast.Ident:
								if _, ok := paramsMap[x.String()]; ok {
									paramsMap[x.String()].used = true
								}
							}
							switch y := rh.Y.(type) {
							case *ast.Ident:
								if _, ok := paramsMap[y.String()]; ok {
									paramsMap[y.String()].used = true
								}
							}
						}
					}
				case *ast.ReturnStmt:
					// TODO: Add all cases of Stmt.

				}
			}

			// Make Issues based on the unused params.
			for name, param := range paramsMap {
				if param.used {
					continue
				}
				issues = append(issues, &Issue{
					Text: fmt.Sprintf("%s is unused in %s", name, dec.Name.String()),
					Pos:  fset.Position(param.ident.Pos()),
				})
			}
		}
		return true
	})
	return issues, nil
}
