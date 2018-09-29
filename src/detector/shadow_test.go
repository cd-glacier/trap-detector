package detector

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestDetectShadow(t *testing.T) {
	tables := []struct {
		code  string
		Found []FoundNode
	}{
		{
			`
package main

func main() {
	x := 0
	y := 0
	z := 0
	if x == 0 {
		x := 1
	} else if false {
		y := 2
	} else {
		z := 3
	}
}
			`,
			[]FoundNode{
				FoundNode{Name: "x"},
				FoundNode{Name: "y"},
				FoundNode{Name: "z"},
			},
		},
		{
			`
package main
func main() {
	x := 0
	y := 0
	if x := 1; true {
	} else if y := 2; true {
	}
}
			`,
			[]FoundNode{
				FoundNode{Name: "x"},
				FoundNode{Name: "y"},
			},
		},
	}

	for _, tt := range tables {
		f, err := parser.ParseFile(token.NewFileSet(), "main.go", tt.code, parser.AllErrors)
		if err != nil {
			t.Fatalf("Failed to parse: %s", err.Error())
		}

		s := &Shadow{}
		s.Detect(f)

		if !reflect.DeepEqual(s.Found, tt.Found) {
			t.Fatalf("Failed to DetectShadow\n actual  :%#v\n expected:%#v", s.Found, tt.Found)
		}
	}
}
