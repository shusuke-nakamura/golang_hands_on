package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"taro", []int{10, 20, 30}}
	hanako := Mydata{
		Name: "Hanako",
		Data: []int{90, 80, 70},
	}
	fmt.Println(taro)
	fmt.Println(hanako)
}
