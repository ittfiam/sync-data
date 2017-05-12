package sync

import (
	"regexp"
	"fmt"
	"strings"
	"time"
)

/**
处理order的特殊规则
init 所有
increment 7天前
 */


type RuleSub struct {
	SourceTb string `json:"source_tb"`
	TargetTB string `json:"target_tb"`
	Columns []string `json:"columns"`
	UpdateTb string `json:"update_tb"`
	UpdateColumn string `json:"update_column"`
	NotNeedTruncate bool `json:"not_need_truncate"`
}

func (subRule *RuleSub) getUpdateColumn() string{

	if subRule.UpdateColumn == ""{
		return "update_time"
	}
	return subRule.UpdateColumn
}

func NewRuleSub(sourceTb string) *RuleSub{
	return &RuleSub{
		SourceTb:sourceTb,
		TargetTB:sourceTb,
		UpdateTb:sourceTb,
		Columns:make([]string,0),
		UpdateColumn:"update_time",
	}
}

type RuleConfig struct {
	SourceDB string `json:"source_db"`
	TargetDB string `json:"target_db"`
	Subs []*RuleSub `json:"sub"`
	All bool `json:"all"`
}


type RuleConfigs []*RuleConfig


func NewRuleConfigs() (*RuleConfigs, error) {

	var r RuleConfigs

	exist, err := AssetExists("runConfig.json")

	fmt.Println(exist)

	if err != nil {
		return nil,err
	}

	if exist {
		err = ReadAssetAsJSON("runConfig.json", &r)
		if err != nil{
			return nil,err
		}
	}

	return &r,nil
}


func (rules *RuleConfigs) GetRule(db string) *RuleConfig{

	for _,rule := range *rules{
		match, _ := regexp.MatchString(rule.SourceDB, db)
		if match{
			return rule
		}
	}
	return nil
}

func (rule *RuleConfig) GetRuleSub(tb string) *RuleSub{
	if rule.Subs == nil{
		rule.Subs = make([]*RuleSub,0)
	}
	if rule.All{
		ruleSub := NewRuleSub(tb)
		rule.Subs = append(rule.Subs,ruleSub)
		return ruleSub
	}
	for _,r := range rule.Subs{
		match, _ := regexp.MatchString(r.SourceTb, tb)
		if match{
			return r
		}
	}

	return nil
}

func (subRule *RuleSub) GetUpdateTable(tables []string) []string{

	today := time.Now()
	result := make([]string,0)
	for _,table := range tables{
		// 只更新前7天的表
		switch subRule.UpdateTb {
		case "$tb_order":
			last := strings.TrimLeft(table,"tb_order_")
			td,err := time.Parse("20060102",last)
			if err != nil{
				break
			}
			d :=today.Sub(td)
			if d.Hours() > 8*24 || d.Hours() < 0{
				break
			}
			result = append(result,table)
		case "$tb_order_goods_detail":
			last := strings.TrimLeft(table,"tb_order_goods_detail_")
			td,err := time.Parse("20060102",last)
			if err != nil{
				break
			}
			d :=today.Sub(td)
			if d.Hours() > 8*24 || d.Hours() < 0{
				break
			}
			result = append(result,table)

		case "$wallet":
			// 只更新2个月的表
			last := strings.TrimLeft(table,"tbl_account_pipeline_")
			td,err := time.Parse("200601",last)
			if err != nil{
				break
			}
			year,month,_ := td.Date()
			cur_year,cur_month,_ := today.Date()
			if year == cur_year{
				d :=cur_month - month
				if  d <2 && d >= 0{
					result = append(result,table)
				}
			}else if cur_year-year== 1{
				if cur_month == 1 && month == 12{
					result = append(result,table)
				}
			}

		default:
			result = append(result,table)
		}
	}

	return result
}