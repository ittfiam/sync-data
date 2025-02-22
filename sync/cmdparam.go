package sync

import (
	"strings"
	"fmt"
)

type CommandParam struct {
	Source string
	Target string
	Prefix string
	Reader string
	Writer string
	Mode int 	// 初始化还是增量描述文件  0:初始化 1:增量
	Path string
	sourceSchema *ConnectScheme
	targetSchema *ConnectScheme
}



func (param *CommandParam) GetSourceSchema()  (*ConnectScheme,error){

	if param.sourceSchema == nil{

		s,err := ParseScheme(param.Source)
		if err != nil{
			return nil,err
		}
		param.sourceSchema = s
	}

	return param.sourceSchema,nil

}

func (param *CommandParam) GetTargetSchema()  (*ConnectScheme,error){

	if param.targetSchema == nil{

		s,err := ParseScheme(param.Target)
		if err != nil{
			return nil,err
		}
		param.targetSchema = s
	}

	return param.targetSchema,nil

}

func (param *CommandParam) GetTransitionMode() string{

	return fmt.Sprintf(
		"%s2%s",
	 strings.TrimRight(param.Reader,"reader"),
	 strings.TrimRight(param.Writer,"writer"),
	)

}

