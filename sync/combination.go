package sync



func Combination(dataxParam *DataXContext,param *CommandParam) *Job{

	switch param.Writer {


	case "mysqlwriter":
		return MysqlMappingMode(dataxParam,param)
	case "hdfswriter":
		return HdfsMappingMode(dataxParam,param)
	default:
		return nil

	}
	return nil
}