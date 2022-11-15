package format

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// Table struct
type Table struct {
	header []string
	order  []int
	rows   [][]string
	cols   int
}

// NewTable returns all the Table row data
func NewTable(h []string) *Table {
	t := &Table{}
	t.SetHeader(h)
	return t
}

// SetHeader sets the header of the Table
func (t *Table) SetHeader(h []string) {
	t.header = h
}

// Append to insert rows to the Table
func (t *Table) Append(row []string) {
	t.rows = append(t.rows, row)
}

// AppendBulk to insert bulk rows to the Table
func (t *Table) AppendBulk(rows [][]string) {
	t.rows = rows
}

// Len for the sort interface
func (t *Table) Len() int {
	return len(t.rows)
}

// Less for the sort interface
func (t *Table) Less(i, j int) bool {
	lastCol := 0
	for _, col := range t.order {
		if t.rows[i][col] != t.rows[j][col] {
			return t.rows[i][col] < t.rows[j][col]
		}
		lastCol = col
	}
	return t.rows[i][lastCol] < t.rows[j][lastCol]
}

// Swap for the sort interface
func (t *Table) Swap(i, j int) {
	t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
}

// OrderBy sets the column order for the Table sort
func (t *Table) OrderBy(order []int) {
	t.order = order
}

// Sort sorts the Table according to the arguments passed to OrderBy.
func (t *Table) Sort() {
	//If no orderBy is set, use by default serial column sort
	if len(t.order) == 0 && t.cols != 0 {
		for idx := 0; idx < t.cols; idx++ {
			t.order = append(t.order, idx)
		}
	}
	sort.Sort(t)
}

// Print Table to stdout using a Table format
func (t *Table) Print() {
	t.Sort()
	Table := tablewriter.NewWriter(os.Stdout)
	if t.header != nil {
		Table.SetHeader(t.header)
	}
	Table.SetAutoWrapText(false)
	Table.SetAutoFormatHeaders(true)
	Table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	Table.SetAlignment(tablewriter.ALIGN_LEFT)
	Table.SetCenterSeparator("")
	Table.SetColumnSeparator("")
	Table.SetRowSeparator("")
	Table.SetHeaderLine(false)
	Table.SetBorder(false)
	Table.SetTablePadding("\t") // pad with tabs
	Table.SetNoWhiteSpace(true)
	Table.AppendBulk(t.rows)
	Table.Render()
}
