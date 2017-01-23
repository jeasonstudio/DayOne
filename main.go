package main

import (
	dayone "DayOne/dayone"
	"fmt"
	"time"
)

func myTicker() {
	m, _ := dayone.Spy()
	dayone.SendEmail(m)
}

func main() {
	SENDTIMES := 1
	ticker := time.NewTicker(time.Hour * 24)
	for _ = range ticker.C {
		fmt.Println("Now Time is ", time.Now().Format("2006-01-02 15:04"), ", sending the ", SENDTIMES, "Email.")
		myTicker()
		SENDTIMES++
	}
}
