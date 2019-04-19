package ghtml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGHtml(t *testing.T) {
	html := New()
	html.
		ChildWithTag(
			"<p>", "This is first paragraph containing table", "</p>",
			ChildWithTag("<table>", "", "</table>").
				ChildWithTag("<tr>", "", "</tr>",
					ChildWithTag("<th>", "First name", "</th>"),
					ChildWithTag("<th>", "Last Name name", "</th>"),
				).
				ChildWithTag("<tr>", "", "</tr>",
					ChildWithTag("<td>", "Faris", "</td>"),
					ChildWithTag("<td>", "Muhammad Syariati", "</td>"),
				),
			ChildWithTag("<b>", "Bold Text", "</b>"),
		).
		ChildWithTag("<p>", "This is second paragraph", "</p>")
	html2 := New()
	html2.
		Child(
			TypeParagraph, "This is first paragraph containing table",
			Child(TypeTable, "").
				Child(TypeTableRow, "",
					Child(TypeTableHeader, "First name"),
					Child(TypeTableHeader, "Last Name name"),
				).
				Child(TypeTableRow, "",
					Child(TypeTableData, "Faris"),
					Child(TypeTableData, "Muhammad Syariati"),
				),
			Child(TypeTextBold, "Bold Text"),
		).
		Child(TypeParagraph, "This is second paragraph")

	assert.Equal(t, html.Scan(), html2.Scan(), "both method equals")
}
