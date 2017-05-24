package sync

import (
	"fmt"
	"strings"
)

var (
	config *TransitionConfig
)

type TColumnTypeItem struct{

	Name []string	`json:"name"`
	Value string	`json:"value"`
}


type TConfigItem struct{

	ColumnType []*TColumnTypeItem	`json:"column_type"`

}


type TransitionConfig map[string]*TConfigItem



func (c *TConfigItem) GetValue(s string) string{

	for _,t := range c.ColumnType{
		for _,n := range t.Name{
			if strings.ToLower(s) == n{
				return t.Value
			}
		}
	}
	return ""
}

func newTransitionConfig() (*TransitionConfig, error) {

	var r TransitionConfig

	exist, err := AssetExists("columnTranstion.json")

	if err != nil {
		return nil,err
	}

	if exist {
		err = ReadAssetAsJSON("columnTranstion.json", &r)
		if err != nil{
			return nil,err
		}
	}

	return &r,nil
}

func GetTransitionConfig() *TransitionConfig{

	if config == nil{
		var err error
		config,err = newTransitionConfig()
		if err != nil{
			fmt.Println("获取转换列配置错误columnTranstion.json")
			fmt.Println(err)
		}
	}
	return config
}

func (t *TransitionConfig) GetTConfigItem(key string) *TConfigItem{


	for name,value := range *t{

		if name == key{
			return value
		}
	}
	return nil
}