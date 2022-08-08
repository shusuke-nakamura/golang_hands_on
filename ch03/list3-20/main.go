package main

import "fmt"

// General is all type data.
type General interface{}

// GData is holding General value.
type GData interface {
	Set(nm string, g General)
	Print()
}

// GDataImple is struct.
type GDataImple struct {
	Name string
	Data General
}

// Set is GDataImple method.
func (gd *GDataImple) Set(nm string, g General) {
	gd.Name = nm
	gd.Data = g
}

func (gd *GDataImple) Print() {
	fmt.Printf("<<%s>>", gd.Name)
	fmt.Println(gd.Data)
}

func main() {
	var data = []GDataImple{}
	data = append(data, GDataImple{"Trao", 123})
	data = append(data, GDataImple{"Hanako", "hello!"})
	data = append(data, GDataImple{"Sachiko", []int{123, 456, 789}})
	for _, ob := range data {
		ob.Print()
	}
}
