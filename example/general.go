package main

import (
	"fmt"
	"goglib"
	"time"
)

func main() {
	const CHECK_INTERVAL1 = 100  // 100 millisecond
	const CHECK_INTERVAL2 = 1000 // 1000 millisecond

	var markTime1 time.Time = time.Now()

	if goglib.CheckElapsedTime(&markTime1, CHECK_INTERVAL1) {
		fmt.Println("100 millisecond was elapsed")
	} else {
		fmt.Println("100 millisecond wasn't elapsed")
	}

	var markTime2 time.Time = time.Now()
	time.Sleep(time.Second * 2)
	if goglib.CheckElapsedTime(&markTime2, CHECK_INTERVAL2) {
		fmt.Println("1000 millisecond was elapsed")
	} else {
		fmt.Println("1000 millisecond wasn't elapsed")
	}
}
