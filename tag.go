package ghtml

import (
	"sync"
)

type ElementType string

const (
	TypeTable       = "table"
	TypeTableRow    = "table-row"
	TypeTableHeader = "table-header"
	TypeTableData   = "table-data"
	TypeParagraph   = "paragraph"
	TypeTextBold    = "bold"
	TypeTextItalian = "italian"
)

var mux sync.Mutex
var tags = map[ElementType]*tag{
	TypeTable:       &tag{start: "<table>", end: "</table>"},
	TypeTableRow:    &tag{start: "<tr>", end: "</tr>"},
	TypeTableHeader: &tag{start: "<th>", end: "</th>"},
	TypeTableData:   &tag{start: "<td>", end: "</td>"},
	TypeParagraph:   &tag{start: "<p>", end: "</p>"},
	TypeTextBold:    &tag{start: "<b>", end: "</b>"},
	TypeTextItalian: &tag{start: "<i>", end: "</i>"},
}

type tag struct {
	start, end string
}

func getTag(e ElementType) (start string, end string) {
	if t, ok := tags[e]; ok {
		start = t.start
		end = t.end
	}
	return
}
