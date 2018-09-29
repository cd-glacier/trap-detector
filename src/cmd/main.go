package main

import (
	"flag"
	"go/parser"
	"go/token"

	"github.com/g-hyoga/trap-detector/src/detector"
	"github.com/g-hyoga/trap-detector/src/finder"
	"github.com/g-hyoga/trap-detector/src/logger"
)

var log = logger.New()

func main() {
	packageName := flag.String("package", "./cmd/", "package which you want to detect trap")
	flag.Parse()

	filenames, err := finder.GetGoFile(*packageName)
	if err != nil {
		log.Error("[main] Failed to finder.GetGoFile: %s", err.Error())
		panic("[main] Failed to finder.GetGoFiles")
	}
	log.Infof("[main] found go files: %+v", filenames)

	for _, filename := range filenames {
		f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.AllErrors)
		if err != nil {
			log.Error("[main] parser.ParseFile: %s", err.Error())
		}

		shadow := &detector.Shadow{}
		shadow.Detect(f)
	}
}
