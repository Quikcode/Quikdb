package core

import (
	"QuikDB/vars"
	"os"
	"path/filepath"
)

func getModules(path string) []string {
	// Make the array empty to store the paths of the files
	var arrayFilesRoutes = make([]string, 0)

	// Get all files of the path
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}

	for _, file := range files {
		extension := filepath.Ext(file.Name())

		if extension == vars.FileExtension {
			arrayFilesRoutes = append(arrayFilesRoutes, filepath.FromSlash(path+"/"+file.Name()))
		}

	}

	return arrayFilesRoutes
}
