package sync

import (
	"strings"

	"encoding/json"

	"fmt"

	"sync-mysql/errors"
)

type Expr interface {
	Eval(value string) bool
}

func (stack *ExprStack) String() string {

	return strings.Join(stack.stack, ".")
}

func (stack *ExprStack) Pop() {
	if len(stack.stack) > 0 {
		stack.stack = stack.stack[0 : len(stack.stack)-1]
	}
}

func NewExprStack(values string) *ExprStack {

	return &ExprStack{
		stack: strings.Split(values, "."),
	}
}

// ==
type EqualExpr struct {
	Value string
}

func (expr *EqualExpr) Eval(value string) bool {
	return expr.Value == value
}

type StartsExpr struct {
	Value string
}

func (expr *StartsExpr) Eval(value string) bool {
	return strings.HasPrefix(value, expr.Value)
}

type InExpr struct {
	Format string
	Begin  Date
	End    Date
	Values map[string]struct{}
}

func (expr InExpr) Eval(value string) bool {

	_, ok := expr.Values[value]

	return ok
}

type ExprStack struct {
	stack []string
}

func (stack *ExprStack) Push(value string) {
	stack.stack = append(stack.stack, value)
}

func condValueTo(source, v interface{}) error {

	bytes, err := json.Marshal(source)

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, v)
}

func ToExpr(cond map[string]interface{}, stack *ExprStack) (expr Expr, err error) {

	if stack == nil {
		stack = NewExprStack("")
	}

	for code, value := range cond {
		switch code {
		case "$eq":
			if !IsString(value) {
				err = errors.NewError("$eq op value must string when parse %s", stack.String())
			}

			expr = &EqualExpr{
				Value: value.(string)}

			return
		case "$start":

			if !IsString(value) {
				err = errors.NewError("$start op value must string when parse %s", stack.String())
			}

			expr = &StartsExpr{
				Value: value.(string)}

			return

		case "$date-before":

			type Input struct {
				Format string `json:"format"`
				End    string `json:"end"`
				Days   int    `json:"days"`
			}

			input := new(Input)
			cond := new(InExpr)

			err = condValueTo(value, input)

			if err != nil {
				err = errors.ToFormatError(
					err,
					"$date-before op value parse fail where parse %s.", stack.String())

				return
			}

			if len(input.Format) == 0 {
				err = errors.NewError(
					"$date-before op value.format not empty where parse %s.", stack.String())
				return
			}

			if input.End == "$now" {
				cond.End = NowDate()
			} else {
				cond.End, err = ParseDate("2016-01-02", input.End)

				if err != nil {
					err = errors.NewError(
						"$date-before parse end fail where parse %s", stack.String())
					return
				}
			}

			if input.Days < 0 || input.Days > 365*1000 {
				err = errors.NewError(
					"$date-before $date-before.days out for range, range [0,365000)")
			}

			cond.Begin = cond.End.Minus(NewDaysDelta(input.Days))
			cond.Values = make(map[string]struct{}, 0)
			DaysIterator(cond.Begin, cond.End, func(date Date) error {

				name := fmt.Sprintf(
					input.Format,
					date.Year(),
					date.Month(),
					date.Day())

				cond.Values[name] = struct{}{}

				return nil
			})

			expr = cond

		default:
			err = errors.NewError("unknow op <%s> when parse %s", stack.String())
		}
	}

	return

}
