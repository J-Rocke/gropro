package main

import (
	"fmt"

	"github.com/J-Rocke/gropro/algorithmus"
	"github.com/J-Rocke/gropro/eingabe"
)

func main() {
	ac, err := eingabe.ParseFromFile("bierkonsum.txt")
	if err != nil {
		fmt.Println("error parsing: ", err)
	}
	l := algorithmus.Loese(ac)
	fmt.Println(l)
}
