package parse

import (
	"QuikDB/src/core/models"
	"fmt"
	"regexp"
	"strings"
)

func Table(input string) (*models.Table, error) {
	// Crea una expresión regular para extraer la información de la tabla
	tableRegex := regexp.MustCompile(`table (\w+) {([^}]*)}`)

	matches := tableRegex.FindStringSubmatch(input)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid table definition")
	}

	// Crea una estructura Table con el nombre de la tabla
	table := &Table{Name: matches[1], Comment: "", Columns: []Column{}}

	// Separa las columnas por líneas
	columnLines := strings.Split(matches[2], "\n")

	// Para cada línea, extrae la información de la columna y la agrega a la tabla
	for _, line := range columnLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		column, err := parseColumn(line)
		if err != nil {
			return nil, err
		}
		table.Columns = append(table.Columns, *column)
	}

	return table, nil
}
