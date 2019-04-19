package ghtml

import (
	"fmt"
)

func New() *Element {
	return &Element{
		childs: make([]*Element, 0),
		tstart: "<html><body>",
		tend:   "</body></html>",
	}
}

func Child(elem ElementType, value string) *Element {
	start, end := getTag(elem)
	return ChildWithTag(start, value, end)
}

func ChildWithTag(tagStart, value, tagEnd string) *Element {
	return &Element{
		childs: make([]*Element, 0),
		tstart: tagStart,
		tend:   tagEnd,
		val:    value,
	}
}

type Element struct {
	childs       []*Element
	tstart, tend string
	val          string
}

func (e *Element) InlineChild(elem ElementType, value string, childs ...*Element) *Element {
	start, end := getTag(elem)
	return e.InlineChildWithTag(start, value, end, childs...)
}

func (e *Element) InlineChildWithTag(tagStart, value, tagEnd string, childs ...*Element) *Element {
	child := &Element{
		childs: make([]*Element, 0),
		tstart: tagStart,
		val:    value,
		tend:   tagEnd,
	}
	child.childs = childs
	e.childs = append(e.childs, child)
	return e
}

func (e *Element) append(cstring string) string {
	return fmt.Sprintf("%s%s%s%s", e.tstart, e.val, cstring, e.tend)
}

func (e *Element) Scan() string {
	var cstring string
	for _, c := range e.childs {
		cstring += c.Scan()
	}
	return e.append(cstring)
}
