package parse

import (
	"QuikDB/src/core/models"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func Table(data string) {

	tableRegex := regexp.MustCompile(`table\s+(\w+)\s*\{((?:\s*\w+\s+\w+\s*\[[^\]]*\])*)\s*\}`)
	columnRegex := regexp.MustCompile(`\s*(\w+)\s+(\w+)\s*\[([^\]]*)\]`)

	tableMatches := tableRegex.FindAllStringSubmatch(data, -1)

	for _, tableMatch := range tableMatches {
		table := models.Table{
			Name: tableMatch[1],
		}

		columnMatches := columnRegex.FindAllStringSubmatch(tableMatch[2], -1)

		var columns []models.Column

		for _, ColumnMatch := range columnMatches {

			columns = append(columns, models.Column{
				Name:        ColumnMatch[0],
				Type:        ColumnMatch[1],
				Constraints: strings.Split(ColumnMatch[2], ","),
			})

			table.Columns = columns
		}

		b, _ := json.Marshal(table)

		fmt.Println(string(b))

	}

}
