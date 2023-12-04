package main

import (
	"log"
)

type myStruct struct {
	FirstName string
}

func

func main() {
	var myVar myStruct
	myVar.FirstName = "Doug"
	myVar2 := myStruct{
		FirstName: "Andrea",
	}
	log.Printf("myVar is set to %s\n", myVar.FirstName)
	log.Printf("myVar2 is set to %s", myVar2.FirstName)
}
