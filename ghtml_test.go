package ghtml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGHtml(t *testing.T) {
	html := New()
	html.
		InlineChildWithTag(
			"<p>", "This is first paragraph containing table", "</p>",
			ChildWithTag("<table>", "", "</table>").
				InlineChildWithTag("<tr>", "", "</tr>",
					ChildWithTag("<th>", "First name", "</th>"),
					ChildWithTag("<th>", "Last Name name", "</th>"),
				).
				InlineChildWithTag("<tr>", "", "</tr>",
					ChildWithTag("<td>", "Faris", "</td>"),
					ChildWithTag("<td>", "Muhammad Syariati", "</td>"),
				),
			ChildWithTag("<b>", "Bold Text", "</b>"),
		).
		InlineChildWithTag("<p>", "This is second paragraph", "</p>")
	html2 := New()
	html2.
		InlineChild(
			TypeParagraph, "This is first paragraph containing table",
			Child(TypeTable, "").
				InlineChild(TypeTableRow, "",
					Child(TypeTableHeader, "First name"),
					Child(TypeTableHeader, "Last Name name"),
				).
				InlineChild(TypeTableRow, "",
					Child(TypeTableData, "Faris"),
					Child(TypeTableData, "Muhammad Syariati"),
				),
			Child(TypeTextBold, "Bold Text"),
		).
		InlineChild(TypeParagraph, "This is second paragraph")

	assert.Equal(t, html.Scan(), html2.Scan(), "both method equals")
}
