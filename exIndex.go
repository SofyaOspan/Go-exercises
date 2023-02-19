package main

import "fmt"

func Index(str string, Find string) int {
	for i := 0; i < len(str)-len(Find); i++ {
		flag := true
		for j := 0; j < len(Find); j++ {
			if str[i+j] != Find[j] {
				flag = false
			}
		}
		if flag == true {
			return i
		}
	}
	return -1
}