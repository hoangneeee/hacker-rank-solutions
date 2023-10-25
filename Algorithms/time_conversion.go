package main

import "time"

func timeConversion(s string) string {
	// Write your code here
	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, _ := time.Parse(layout1, s)

	// fmt.Println(t.Format(layout1))
	// fmt.Println(t.Format(layout2))
	return t.Format(layout2)
}
