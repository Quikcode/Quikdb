package core

import (
	"QuikDB/src/core/lib"
	"QuikDB/src/core/parse"
)

func analyzeModule(path string) {
	data, err := lib.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	parse.Table(data)
}
