package utils

import (
	"fmt"
	"time"
)


func ParseTime(s string) time.Time  {
	fmt.Println("=====================")
	fmt.Printf(" in time with %v \n", s)
	fmt.Println("====================")
	fmt.Printf("%v\n", s)
	dateRead, err := time.Parse("2006/01/02", s)
	if err != nil {
		fmt.Printf("error converting time\n: %v", err)
		return time.Time{}
	}
	return dateRead
}