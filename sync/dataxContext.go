package sync



type DataXContext struct {

	DbName string
	SourceTable []string
	Sql string
	Rule *RuleConfig
	SubRule *RuleSub
	Table *Table
}

func NewDataxContext() *DataXContext{

	return &DataXContext{
		SourceTable:make([]string,0),
	}
}