package main

import "fmt"

type General interface{}

type GData interface {
	Set(nm string, g General)
	Print()
}

type GDataImple struct {
	Name string
	Data General
}

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
	data = append(data, GDataImple{"Taro", 123})
	data = append(data, GDataImple{"Hanako", "hello"})
	data = append(data, GDataImple{"Sachiko", []int{123, 456, 789}})
	for _, ob := range data {
		ob.Print()
	}
}
