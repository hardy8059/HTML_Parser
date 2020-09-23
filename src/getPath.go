package src

import (
	"path/filepath"
	"runtime"
)

// ProjectPaths represents paths to internal project files relative to root folder...
type ProjectPaths struct {
	pathToTemplates string
}

func getBasePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}

func initialisePaths() *ProjectPaths {
	var p ProjectPaths
	basePath := getBasePath()
	p.pathToTemplates = filepath.Join(basePath, "data", "templates")
	return &p
}
