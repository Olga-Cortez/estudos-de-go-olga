package main

import "fmt"

func parte2(){
	for i:=1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println("Pin")
		} else if i % 5 == 0 {
			fmt.Println("Pan")
		}
	}
}