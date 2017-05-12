package sync




type DataXContext struct {

	DbName string
	SourceTable []string
	Sql string
	Rule *RuleConfig
	SubRule *RuleSub
	SourceScheme *ConnectScheme
	TargetScheme *ConnectScheme
}

func NewDataxContext() *DataXContext{

	return &DataXContext{
		SourceTable:make([]string,0),
	}
}