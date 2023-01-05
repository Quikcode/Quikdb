package core

import (
	"QuikDB/src/core/models"
	"regexp"
)

func GetImports(input []byte) []models.Import {

	regex := regexp.MustCompile(`import (\w+) from "(.+)"`)

	matches := regex.FindAllStringSubmatch(string(input), -1)

	imports := make([]models.Import, 0)

	for _, match := range matches {
		imports = append(imports, models.Import{
			Name: match[1],
			Path: match[2],
		})
	}

	return imports

}
