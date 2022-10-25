package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
)

func main(){
	layout := "2006-01-02 15:04:05"
	fmt.Print("Enter your text: ")
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	reader = strings.TrimRight(reader, "\n")
	s := strings.Split(reader, ",")
	fmt.Println(s)
	s1, _ := time.Parse(layout, s[0])
	s2, _ := time.Parse(layout, s[1])
	if s1.After(s2){
		fmt.Println(s1.Sub(s2))
	} else {
		fmt.Println(s2.Sub(s1))
	}
	// fmt.Println(s)
}