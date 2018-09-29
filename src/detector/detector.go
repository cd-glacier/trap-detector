package detector

import "go/ast"

type Detector interface {
	Detect(*ast.Node) bool
}

type FoundNode struct {
	Name string
	Line int
}
