package core

import (
	"QuikDB/lib/files"
	"QuikDB/src/core/models"
)

var (
	Tables     []models.Table
	References []models.Reference
)

func Run() {
	read, _ := files.Read("E:\\Proyectos\\github.com\\Quikcode\\Quikdb\\test\\test_table.qdb")

}
