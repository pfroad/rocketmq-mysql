package schema

type Table struct {
	SchemaName string
	Name       string
	ColNames   []string
}

func (self *Table) AddCol(col string) {
	self.ColNames = append(self.ColNames, col)
}
