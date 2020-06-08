package main

import (	
	"fmt"
	"time"
)

func main(){
	today := time.Now().Format("2006/01/02")
	tmr := time.Now().Add(24*time.Hour).Format("2006/01/02")

	fmt.Println(today)
	fmt.Println(tmr)
	//fmt.Println(today + 1)

}

