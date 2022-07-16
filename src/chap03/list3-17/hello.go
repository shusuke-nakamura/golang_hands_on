package main

import "fmt"

type Data interface {
	Initial(name string, data []int)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]\n", md.Name)
}

func main() {
	var ob Data = new(Mydata)
	ob.Initial("Sachiko", []int{55, 66, 77})
	// errror
	// ob.Check()
}
