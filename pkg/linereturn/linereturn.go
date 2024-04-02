package linereturn

import (
	"flag"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const (
	linterName = "linereturn"
	linterDoc  = `Linter requires a new line before return statements if the previous line in a block is a real statement.`
)

var blockSize int

// NewAnalyzer returns a new linereturn analyzer.
func NewAnalyzer() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name: linterName,
		Doc:  linterDoc,
		Run:  run,
	}
	a.Flags.Init("linereturn", flag.ExitOnError)
	a.Flags.IntVar(&blockSize, "block-size", 1, "set block size that is still ok")

	return a
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			switch c := node.(type) {
			case *ast.CaseClause:
				inspectBlock(pass, c.Body)
			case *ast.CommClause:
				inspectBlock(pass, c.Body)
			case *ast.BlockStmt:
				inspectBlock(pass, c.List)
			}
			return true
		})
	}
	return nil, nil
}

func inspectBlock(pass *analysis.Pass, block []ast.Stmt) {
	for i, stmt := range block {
		switch stmt.(type) {
		case *ast.ReturnStmt:
			if m := isNok(pass, block, i); m != "" {
				pass.Report(analysis.Diagnostic{
					Pos:     stmt.Pos(),
					Message: m,
					SuggestedFixes: []analysis.SuggestedFix{
						{
							TextEdits: []analysis.TextEdit{
								{
									Pos:     stmt.Pos(),
									NewText: []byte("\n"),
									End:     stmt.Pos(),
								},
							},
						},
					},
				})
			}
		}
	}
}

func isNok(pass *analysis.Pass, block []ast.Stmt, i int) string {
	stmt := block[i]
	if i == 0 || line(pass, stmt.Pos())-line(pass, block[0].Pos()) < blockSize {
		return ""
	}
	ret := true
	switch s := block[i-1].(type) {
	case *ast.BranchStmt, *ast.ReturnStmt:
		ret = false
	case *ast.BlockStmt:
		ret = false
	case *ast.IfStmt, *ast.SwitchStmt, *ast.SelectStmt, *ast.ForStmt, *ast.RangeStmt, *ast.TypeSwitchStmt:
		ret = false
	case *ast.AssignStmt:
		if len(s.Rhs) <= 1 {
			switch s.Rhs[0].(type) {
			case *ast.CompositeLit, *ast.UnaryExpr, *ast.CallExpr:
				if line(pass, block[i-1].End())-line(pass, block[i-1].Pos()) > 1 {
					ret = false
				}
			}
		}
	case *ast.ExprStmt, *ast.GoStmt, *ast.DeferStmt:
		if line(pass, block[i-1].End())-line(pass, block[i-1].Pos()) > 1 {
			ret = false
		}
	default:
	}
	if ret && line(pass, stmt.Pos())-line(pass, block[i-1].End()) <= 1 {
		return "no blank line before"
	}
	if !ret && line(pass, stmt.Pos())-line(pass, block[i-1].End()) > 1 {
		return "no blank line needed"
	}
	return ""
}

func line(pass *analysis.Pass, pos token.Pos) int {
	return pass.Fset.Position(pos).Line
}
