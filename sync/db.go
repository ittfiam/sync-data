package sync

type DB struct {
	Name   string
	Length int
	Index  int
	Tables map[string]*Table
}
