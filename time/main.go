package main

import (
	"fmt"
	"time"
)

func main(){
	var s string
	fmt.Scanln(&s)
	v, _ := time.Parse(time.RFC3339, s)
	new1 := v.Format(time.UnixDate)
	fmt.Println(new1)
}