package domain

import (
	"net/http"
	"text/template"
)

type MultiplicationTable struct {
	TableCategory int
	TableHeader   []int
	TableRow      [][]int
}

type Data struct {
	Title  string // Page title
	Tables []MultiplicationTable
}

func GenerateMultiplicationTable(n int) []MultiplicationTable {
	tables := make([]MultiplicationTable, 0)

	for base := 1; base <= n; base++ {
		tableRow := make([][]int, base)
		tableHeader := make([]int, 0)
		for i := 1; i <= base; i++ {
			tableHeader = append(tableHeader, i)
			tableRow[i-1] = append(tableRow[i-1], i)
			for j := 1; j <= base; j++ {
				tableRow[i-1] = append(tableRow[i-1], i*j)
			}
		}

		table := MultiplicationTable{
			TableCategory: base,
			TableHeader:   tableHeader,
			TableRow:      tableRow,
		}

		tables = append(tables, table)
	}

	return tables
}

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content Type", "text/html")

	// Initialize new template
	indexTemplate := template.New("indexTemplate")

	// Create template called "doc" in memory using IndexPages html
	indexTemplate.New("doc").Parse(IndexPages)

	// Generate multiplication table
	data := &Data{
		Title:  "Go multiplication table",
		Tables: GenerateMultiplicationTable(99),
	}

	// Look for created template called "doc" in memory, write to rw
	indexTemplate.Lookup("doc").Execute(rw, data)
}
