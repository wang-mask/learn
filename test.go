package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// mm, _ := time.ParseDuration("29h")
	// next := now.Add(mm)
	// fmt.Println(now, next)

	// fmt.Println(next.Sub(now).Hours())
	timeStr := "2018-01-01 11:20:00"
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	now := time.Now()
	fmt.Println("timeStr:", now.Sub(t))
}
