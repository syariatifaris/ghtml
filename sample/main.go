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
		Child(
			ghtml.TypeParagraph, "This is first paragraph containing table",
			ghtml.Child(ghtml.TypeTable, "").
				Child(ghtml.TypeTableRow, "",
					ghtml.Child(ghtml.TypeTableHeader, "First name"),
					ghtml.Child(ghtml.TypeTableHeader, "Last Name name"),
				).
				Child(ghtml.TypeTableRow, "",
					ghtml.Child(ghtml.TypeTableData, "Faris"),
					ghtml.Child(ghtml.TypeTableData, "Muhammad Syariati"),
				),
			ghtml.Child(ghtml.TypeTextBold, "Bold Text"),
		).
		Child(ghtml.TypeParagraph, "This is second paragraph")
	return html.Scan()
}

func withTag() string {
	html := ghtml.New()
	html.
		ChildWithTag(
			"<p>", "This is first paragraph containing table", "</p>",
			ghtml.ChildWithTag("<table>", "", "</table>").
				ChildWithTag("<tr>", "", "</tr>",
					ghtml.ChildWithTag("<th>", "First name", "</th>"),
					ghtml.ChildWithTag("<th>", "Last Name name", "</th>"),
				).
				ChildWithTag("<tr>", "", "</tr>",
					ghtml.ChildWithTag("<td>", "Faris", "</td>"),
					ghtml.ChildWithTag("<td>", "Muhammad Syariati", "</td>"),
				),
			ghtml.ChildWithTag("<b>", "Bold Text", "</b>"),
		).
		ChildWithTag("<p>", "This is second paragraph", "</p>")
	return html.Scan()
}
