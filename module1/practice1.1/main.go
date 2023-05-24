package main

import (
	"fmt"
)

func main() {

	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("数组初始值如下：")
	fmt.Print("[")
	for index, value := range myArray {
		if index == 4 {
			fmt.Print(value)
		} else {
			fmt.Print(value + ",")
		}
		if value == "stupid" {
			myArray[index] = "smart"
		} else if value == "weak" {
			myArray[index] = "strong"
		}
	}
	fmt.Println("]")
	fmt.Println("数组修改后的值如下：")
	fmt.Print("[")
	for index, value := range myArray {
		if index == 4 {
			fmt.Print(value)
		} else {
			fmt.Print(value + ",")
		}
	}
	fmt.Println("]")
}
