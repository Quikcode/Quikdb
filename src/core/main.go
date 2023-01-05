package core

import (
	"QuikDB/src/core/models"
)

var (
	Tables     []models.Table
	References []models.Reference
	Modules    []string
)

func Run() {
	// Get all files of the route
	Modules = getModules("E:\\Proyectos\\github.com\\Quikcode\\Quikdb\\test\\database")

	// Analyze modules to find Tables and References
	for _, path := range Modules {
		analyzeModule(path)
	}

}
