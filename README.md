# GHTML - Backend Generated HTML Text

## Sample Usage

Without Tag
```
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
fmt.Println(html.Scan())
```

With Tag:
```
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
fmt.Println(html.Scan())
```

Both will produce:
```
<html><body><p>This is first paragraph containing table<table><tr><th>First name</th><th>Last Name name</th></tr><tr><td>Faris</td><td>Muhammad Syariati</td></tr></table><b>Bold Text</b></p><p>This is second paragraph</p></body></html>
```