package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	layout := "2006-01-02 15:04:05"
	fmt.Print("Enter your text: ")
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	reader = strings.TrimRight(reader, "\n")
	s, _ := time.Parse(layout, reader)
	if s.Hour() > 13 {
		newtime := s.AddDate(0, 0, 1)
		fmt.Println(newtime.Format(layout))
	} else {
		fmt.Println(s.Format(layout))
	}
}