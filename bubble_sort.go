package main

import(
	"fmt"
)

func main(){
	a := []int{1, 5, 8, 9, 1, 4, 6, 7, 2, 5, 9, 1, 5, 10, 6}
	sort(a)
	fmt.Println(a)
}

func sort(a []int){
	i := 0
	for i < len(a){
		indexOne := 0
		indexTwo := 1
		for indexTwo < len(a){
			varOne := a[indexOne]
			varTwo := a[indexTwo]
		
			if varTwo < varOne{
				a[indexTwo] = varOne
				a[indexOne] = varTwo
			}		
			indexOne++
			indexTwo++
		}
		i++
	}
}
