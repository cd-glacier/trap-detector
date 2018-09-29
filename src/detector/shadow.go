package detector

import (
	"go/ast"

	"github.com/g-hyoga/trap-detector/src/logger"
)

type Shadow struct {
	VarNodes []ast.Ident
	Found    []FoundNode
}

var log = logger.New()

func (s *Shadow) Detect(file *ast.File) {
	for _, decl := range file.Decls {
		s.detect(decl)
	}
}

func (s *Shadow) detect(decl ast.Decl) {
	switch n := decl.(type) {
	case *ast.FuncDecl:
		log.Infof("[shadow] found %s FuncDel.", n.Name.Name)
		// initialize
		s.VarNodes = []ast.Ident{}
		s.detectBlockStmt(n.Body)
	}
}

func (s *Shadow) detectBlockStmt(block *ast.BlockStmt) {
	for _, stmt := range block.List {
		s.detectStmt(stmt)
	}
}

func (s *Shadow) detectStmt(statement ast.Stmt) {
	switch stmt := statement.(type) {
	case *ast.IfStmt:
		s.detectIf(*stmt)
	case *ast.ForStmt:
		s.detectFor(*stmt)
	case *ast.AssignStmt:
		s.detectAssignStmt(*stmt)
	case *ast.BlockStmt: // else
		s.detectBlockStmt(stmt)
	}
}

func (s *Shadow) detectAssignStmt(stmt ast.AssignStmt) {
	for _, hs := range stmt.Lhs {
		i := hs.(*ast.Ident)
		if s.contains(i) {
			log.Infof("[shadow] found shadow!!: %s is duplicated.", i.Name)
			foundNode := FoundNode{Name: i.Name}
			s.Found = append(s.Found, foundNode)
		} else {
			log.Infof("[shadow] add variable: %s", i.Name)
			s.VarNodes = append(s.VarNodes, *i)
		}
	}
}

func (s *Shadow) detectIf(stmt ast.IfStmt) {
	if stmt.Init != nil {
		s.detectStmt(stmt.Init)
	}
	s.detectBlockStmt(stmt.Body)
	s.detectStmt(stmt.Else)
}

func (s *Shadow) detectFor(stmt ast.ForStmt) {
	s.detectBlockStmt(stmt.Body)
}

func (s *Shadow) detectExpr(expr ast.Expr) {
	switch e := expr.(type) {
	case *ast.Ident:
		if e.Obj != nil {
			stmt := e.Obj.Decl.(*ast.AssignStmt)
			s.detectAssignStmt(*stmt)
		}
	}
}

func (s *Shadow) detectIdent(expr ast.Expr) {

}

func (s *Shadow) contains(ident *ast.Ident) bool {
	for _, v := range s.VarNodes {
		if v.Name == ident.Name && v.Pos() < ident.Pos() {
			return true
		}
	}
	return false
}
