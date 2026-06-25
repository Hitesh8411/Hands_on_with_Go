package main

import (
	"fmt"
	"time"
)

// Go alng date and time format :

/*
2006-01-02 15:04:05
2006 : year
01 : Month
02 : date

15 : 24-hour format
04 : minute part of time
05 : second part of time
*/
func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	formatted := currentTime.Format("02-01-2006, Monday, 15:04:05") //work only this format
	amTime := currentTime.Format("3:04 PM")
	fmt.Println("Formatted time: ",formatted,amTime)

	layout_str := "2006-01-02"
	dataStr := "2023-11-25"
	formatted_time,_:= time.Parse(layout_str, dataStr)
	fmt.Println("Formatted time:", formatted_time)
    
	// add 1 more day in currentTime
    new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new_date time:", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date time: ", formatted_new_date)
}
