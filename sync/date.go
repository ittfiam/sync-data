package sync

import (
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

type DateDelta struct {
	years  int
	months int
	days   int
}

func NewDaysDelta(days int) DateDelta {
	return DateDelta{
		days: days,
	}
}

func (date Date) Year() int {
	return date.year
}

func (date Date) Month() int {
	return date.month
}

func (date Date) Day() int {
	return date.day
}

func (date Date) Minus(delta DateDelta) Date {

	t := date.toTime()
	t = t.AddDate(
		-1*delta.years,
		-1*delta.months,
		-1*delta.days,
	)
	return NewDateFromTime(t)
}

func (date Date) AddDays(days int) Date {

	t := date.toTime()
	t = t.AddDate(0, 0, days)

	return NewDateFromTime(t)
}

func NowDate() Date {

	now := time.Now()

	return NewDateFromTime(now)

}

func DaysIterator(begin, end Date, cb func(date Date) error) error {

	for !begin.Before(end) {

		err := cb(begin)

		if err != nil {
			return err
		}

		begin = begin.AddDays(1)
	}

	return nil

}

func (date Date) Equal(o Date) bool {

	return date.year == o.year && date.month == o.month && date.day == o.day
}

func (date Date) ToNumber() int {

	return date.year*10000 + date.month*100 + date.day
}

func (date Date) toTime() time.Time {

	return time.Date(
		date.year,
		time.Month(date.month),
		date.day,
		0,
		0,
		0,
		0,
		time.UTC,
	)
}

func (date Date) Before(o Date) bool {
	return date.ToNumber() > o.ToNumber()
}

func NewDateFromTime(t time.Time) Date {

	return Date{
		year:  t.Year(),
		month: int(t.Month()),
		day:   t.Day(),
	}
}

func ParseDate(layout, value string) (Date, error) {

	t, err := time.Parse(layout, value)

	if err != nil {
		return Date{}, err
	}

	return NewDateFromTime(t), nil
}
