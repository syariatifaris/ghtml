package main

import (
	"fmt"
	"github.com/syariatifaris/ghtml"
)

func main() {
	fmt.Println(withoutTag())
	fmt.Println(withTag())
}

func withoutTag() string {
	html := ghtml.New()
	html.
		InlineChild(
			ghtml.TypeParagraph, "This is first paragraph containing table",
			ghtml.Child(ghtml.TypeTable, "").
				InlineChild(ghtml.TypeTableRow, "",
					ghtml.Child(ghtml.TypeTableHeader, "First name"),
					ghtml.Child(ghtml.TypeTableHeader, "Last Name name"),
				).
				InlineChild(ghtml.TypeTableRow, "",
					ghtml.Child(ghtml.TypeTableData, "Faris"),
					ghtml.Child(ghtml.TypeTableData, "Muhammad Syariati"),
				),
			ghtml.Child(ghtml.TypeTextBold, "Bold Text"),
		).
		InlineChild(ghtml.TypeParagraph, "This is second paragraph")
	return html.Scan()
}

func withTag() string {
	html := ghtml.New()
	html.
		InlineChildWithTag(
			"<p>", "This is first paragraph containing table", "</p>",
			ghtml.ChildWithTag("<table>", "", "</table>").
				InlineChildWithTag("<tr>", "", "</tr>",
					ghtml.ChildWithTag("<th>", "First name", "</th>"),
					ghtml.ChildWithTag("<th>", "Last Name name", "</th>"),
				).
				InlineChildWithTag("<tr>", "", "</tr>",
					ghtml.ChildWithTag("<td>", "Faris", "</td>"),
					ghtml.ChildWithTag("<td>", "Muhammad Syariati", "</td>"),
				),
			ghtml.ChildWithTag("<b>", "Bold Text", "</b>"),
		).
		InlineChildWithTag("<p>", "This is second paragraph", "</p>")
	return html.Scan()
}
