package sync



func Combination(dataxParam *DataXContext,param *CommandParam) *Job{

	switch param.Writer {


	case "mysqlwriter":
		return MysqlCombinationInit(dataxParam,param)
	case "hdfswriter":
		return HdfsCombinationInit(dataxParam,param)
	default:
		return nil

	}
	return nil
}