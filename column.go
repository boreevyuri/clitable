package clitable

import "unicode/utf8"

var (
	defaultHeaderStyle = &ColumnStyle{
		Align: ColumnAlignCenter,
		VerticalAlign: ColumnVerticalAlignMiddle,
	}
	defaultBodyStyle = &ColumnStyle{
		Align: ColumnAlignLeft,
		VerticalAlign: ColumnVerticalAlignTop,
	}
)

type ColumnAlign int

const (
	ColumnAlignLeft   = iota
	ColumnAlignCenter
	ColumnAlignRight
)

type ColumnVerticalAlign int

const (
	ColumnVerticalAlignTop    = iota
	ColumnVerticalAlignMiddle
	ColumnVerticalAlignBottom
)

type ColumnStyle struct {
	Align         ColumnAlign
	VerticalAlign ColumnVerticalAlign
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int
	PaddingLeft   int
}

type Column struct {
	width       int
	HeaderStyle *ColumnStyle
	BodyStyle   *ColumnStyle
}

func NewColumn(name string) *Column {
	return &Column{
		width      : utf8.RuneCountInString(name),
		HeaderStyle: defaultHeaderStyle,
		BodyStyle  : defaultBodyStyle,
	}
}

func (this *Column) getStyleByRow(row *Row) *ColumnStyle {
	if row.isHeader {
		return this.HeaderStyle
	} else {
		return this.BodyStyle
	}
}
