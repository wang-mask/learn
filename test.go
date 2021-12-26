package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	mm, _ := time.ParseDuration("29h")
	next := now.Add(mm)
	fmt.Println(now, next)

	fmt.Println(next.Sub(now).Hours())
}
