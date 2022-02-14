package examples

import (
	"fmt"
	"time"
)

func getDateTuple(t time.Time) (int, time.Month, int) {
	return t.Day(), t.Month(), t.Year()
}

func printTime(t time.Time)  {
	fmt.Printf("[%02d:%02d:%02d] (%02d/%02d/%4d)\n", t.Hour(), t.Minute(), t.Second(), t.Day(), t.Month(), t.Year())
}

func PrintCurrentTime() {
	t := time.Now()
	printTime(t)
}

func PrintTimeInFuture(dayOffset int)  {
	t := time.Now()
	t2 := t.Add(time.Hour * 24 * time.Duration(dayOffset))
	tday, tmonth, tyear := getDateTuple(t)
	t2day, t2month, t2year := getDateTuple(t2)
	fmt.Printf("In %02d days from (%02d/%02d/%4d) is going to be ", dayOffset, tday, tmonth, tyear)
	fmt.Printf("[%02d/%02d/%4d]\n", t2day, t2month, t2year)
}

func PrintTimeOfDead() {
	t := time.Now()
	years := 2040 - t.Year()
	t = t.Add(time.Hour * 24 * 365 * time.Duration(years))
	fmt.Print("Your time of dead is: ")
	printTime(t)
}