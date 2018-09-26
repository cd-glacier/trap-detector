package detector

import "go/ast"

type Detector interface {
	Detect(*ast.Node) bool
}
