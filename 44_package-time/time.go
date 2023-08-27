package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	nowHour := now.Add(time.Hour)
	now2Hour := now.Add(time.Hour * 2)

	fmt.Println(now)
	fmt.Println(nowHour)
	fmt.Println(now2Hour)
}
