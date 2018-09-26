package detector

import "go/ast"

type Shadow struct {
	Line int
}

func (s *Shadow) Detect(node *ast.Node) bool {
	switch n := node.(type) {
	case *ast.IfStmt:
		pp.Println(n)
	case *ast.ForStmt:
		pp.Println(n)
	}
}
