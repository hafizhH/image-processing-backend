package tests

import (
	"os"
	"path"
	"runtime"
)

func init() {
	// Change test working dir to project root
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)

	if err != nil {
		panic(err)
	}
}
