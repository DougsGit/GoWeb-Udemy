package main

import "fmt"

func main() {
	firstLine :="Once upon a midnight dreary"
	for i, letter:=range firstLine {
		fmt.Printf("the %d letter is %v\n",i,string(letter))
	}
}
