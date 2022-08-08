package main

import (
	"fmt"
	"reflect"
)

type General interface{}

type GData interface {
	Set(nm string, g General) GData
	Print()
}

type NData struct {
	Name string
	Data int
}

func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	if reflect.TypeOf(g).Kind() == reflect.Int {
		nd.Data = g.(int)
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

type SData struct {
	Name string
	Data string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	if reflect.TypeOf(g).Kind() == reflect.String {
		sd.Data = g.(string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func main() {
	var data = []GData{}
	data = append(data, new(NData).Set("Taro", 123))
	data = append(data, new(SData).Set("Jiro", "hello!"))
	data = append(data, new(NData).Set("Hanako", "98700"))
	data = append(data, new(SData).Set("Sachiko", []string{"happly?"}))

	for _, ob := range data {
		ob.Print()
	}
}
