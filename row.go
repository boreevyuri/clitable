package clitable

type Row struct {
	height   int
	cells    []*Cell
	isHeader bool
}

func NewRow() *Row {
	return &Row{
		height: 1,
		cells:  make([]*Cell, 0),
	}
}
