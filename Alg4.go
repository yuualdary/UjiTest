package main

import (
	"fmt"
)

func Range(n int) {

	lamp := 0
	flag := 1
	for i := 1; i <= n; i++ {

		if i %1 == 0 &&flag == 1 {

			lamp = lamp + n
			flag = flag + 1
		} else if i%2 == 0 && flag == 2 {

			lamp = lamp + n/2
			flag = flag + 1


		} else if i%3 == 0 && flag == 3 {

			lamp = lamp + n/3

			flag = 1
		  	//fmt.Println(flag)

		}

	}
	fmt.Println(lamp)
	//3211
}

func main() {
	Range(100)
	fmt.Println("Hello, playground")
}
