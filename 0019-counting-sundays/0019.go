package main

import (
	"fmt"
	"time"
)

const Monday = "Monday"
const Tuesday = "Tuesday"
const Wednesday = "Wednesday"
const Thursday = "Thursday"
const Friday = "Friday"
const Saturday = "Saturday"
const Sunday = "Sunday"

var DayMap = map[string]string{
	Monday:    Tuesday,
	Tuesday:   Wednesday,
	Wednesday: Thursday,
	Thursday:  Friday,
	Friday:    Saturday,
	Saturday:  Sunday,
	Sunday:    Monday,
}

const Jan = "January"
const Feb = "February"
const Mar = "March"
const Apr = "April"
const May = "May"
const Jun = "June"
const Jul = "July"
const Aug = "August"
const Sep = "September"
const Oct = "October"
const Nov = "November"
const Dec = "December"

var MonthMap = map[string]string{
	Jan: Feb,
	Feb: Mar,
	Mar: Apr,
	Apr: May,
	May: Jun,
	Jun: Jul,
	Jul: Aug,
	Aug: Sep,
	Sep: Oct,
	Oct: Nov,
	Nov: Dec,
	Dec: Jan,
}

var MonthLen = map[string]int{
	Jan: 31,
	Mar: 31,
	Apr: 30,
	May: 31,
	Jun: 30,
	Jul: 31,
	Aug: 31,
	Sep: 30,
	Oct: 31,
	Nov: 30,
	Dec: 31,
}

type Year struct {
	year int
}

func (y *Year) Next() {
	y.year += 1
}

type Month struct {
	Name string
}

func (m *Month) Next() {
	m.Name = MonthMap[m.Name]
}

func (m *Month) LastDay(year, dayNum int) bool {
	if m.Name != Feb {
		return dayNum == MonthLen[m.Name]
	} else {
		lastDay := 0
		if year%4 != 0 {
			lastDay = 28
		} else if year%400 == 0 {
			lastDay = 29
		} else if year%100 == 0 {
			lastDay = 28
		} else {
			lastDay = 29
		}
		return dayNum == lastDay
	}

}

type Day struct {
	Name   string
	Number int
}

func (d *Day) Next(lastDay bool) {
	if lastDay {
		d.Name = DayMap[d.Name]
		d.Number = 1
	} else {
		d.Name = DayMap[d.Name]
		d.Number += 1
	}
}

type Date struct {
	year  Year
	month Month
	day   Day
}

func (d *Date) NextDay() {
	if d.month.Name == Dec && d.month.LastDay(d.year.year, d.day.Number) {
		d.year.Next()
	}

	if d.month.LastDay(d.year.year, d.day.Number) {
		d.month.Next()
		d.day.Next(true)
	} else {
		d.day.Next(false)
	}
}

func NewDate(year int, month string, dayName string, dayNum int) Date {
	d := Date{}
	d.year = Year{year}
	d.month = Month{month}
	d.day = Day{dayName, dayNum}

	return d
}

func doit() {
	sunCount := 0
	d := NewDate(1900, Jan, Monday, 1)

	for d.year.year < 2001 {
		if d.year.year > 1900 && d.day.Number == 1 && d.day.Name == Sunday {
			sunCount += 1
		}
		d.NextDay()
	}

	fmt.Println("Total Sundays is", sunCount)
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	doit()
}
