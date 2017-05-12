package sync

import (
	"fmt"
	"testing"
)

func Test_date(a *testing.T) {

	begin, _ := ParseDate("2006-01-02", "2017-10-20")

	end, _ := ParseDate("2006-01-02", "2017-12-21")

	DaysIterator(begin, end, func(date Date) error {

		fmt.Println(date.ToNumber())
		return nil
	})
}

func Test_Delta(t *testing.T) {

	end, _ := ParseDate("2006-01-02", "2017-10-20")
	begin := end.Minus(NewDaysDelta(30))

	DaysIterator(begin, end, func(date Date) error {
		fmt.Println(date.ToNumber())
		return nil
	})
}
