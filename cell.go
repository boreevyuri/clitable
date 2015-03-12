package clitable

import (
	"fmt"
	"unicode/utf8"
)

type Cell struct {
	data     string
	parts    []string
	partsLen int
	width    int
}

func NewCell(data interface {}) *Cell {
	str := fmt.Sprintf("%v", data)
	return &Cell{
		data: str,
		width: utf8.RuneCountInString(str),
	}
}
