package finder

import (
	"os"
	"strings"

	"github.com/g-hyoga/trap-detector/src/logger"
)

var log = logger.New()

func isGoFile(filename string) bool {
	isGoFile := filename[len(filename)-3:] == ".go"
	isTestFile := strings.Contains(filename, "test.")
	return isGoFile && !isTestFile
}

func joinPath(path ...string) string {
	return strings.Replace(strings.Join(path, "/"), "//", "/", -1)
}

func GetGoFile(dir string) ([]string, error) {
	foundFiles := []string{}

	directory, err := os.Open(dir)
	if err != nil {
		return foundFiles, err
	}

	objects, err := directory.Readdir(-1)
	if err != nil {
		return foundFiles, err
	}

	for _, obj := range objects {
		if !obj.IsDir() && isGoFile(obj.Name()) {
			foundFiles = append(foundFiles, joinPath(dir, obj.Name()))
		}
	}

	log.Debugf("[finder] found mutate files: %#v", foundFiles)

	return foundFiles, nil
}
