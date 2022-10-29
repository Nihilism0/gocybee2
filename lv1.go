package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d\n", &n)
	a := make(map[int]int, n)
	str := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &str[i])
	}
	for _, val := range str {
		judge := false
		for k, _ := range a {
			if k == val {
				judge = true
			}
		}
		if judge == false {
			a[val] = 1
		} else {
			a[val]++
		}
	}
	for k, v := range a {
		fmt.Printf("%d有%d个\n", k, v)
	}
}
