package main

import "fmt"

func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}

	for index, value := range arr {
		if value == "stupid" {
			arr[index] = "smart"
		}
		if value == "week" {
			arr[index] = "strong"
		}
		fmt.Println(value)
	}

	for i := 0; i < len(arr); i++ {
		if i == 2 {
			arr[i] = "smart"
		}
		if i == 4 {
			arr[i] = "strong"
		}
		fmt.Println(arr[i])
	}

	fmt.Println(arr)
}
