package models

type Table struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Columns []Column `json:"columns"`
}
