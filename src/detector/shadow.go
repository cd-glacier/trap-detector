package detector

import (
	"go/ast"

	"github.com/g-hyoga/trap-detector/src/logger"
	"github.com/k0kubun/pp"
)

type Shadow struct {
	VarNodes []ast.Ident
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
		log.Infof("found %s FuncDel.", n.Name.Name)
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
		for _, hs := range stmt.Lhs {
			i := hs.(*ast.Ident)
			if s.contains(i) {
				pp.Println("found!!")
			} else {
				s.VarNodes = append(s.VarNodes, *i)
			}
		}

		log.Infof("[shadow] add variable: %#v", s.VarNodes)
	}
}

func (s *Shadow) detectIf(stmt ast.IfStmt) {
	s.detectBlockStmt(stmt.Body)
}

func (s *Shadow) detectFor(stmt ast.ForStmt) {
	s.detectBlockStmt(stmt.Body)
}

func (s *Shadow) contains(ident *ast.Ident) bool {
	for _, v := range s.VarNodes {
		if v.Name == ident.Name && v.Pos() < ident.Pos() {
			return true
		}
	}
	return false
}
