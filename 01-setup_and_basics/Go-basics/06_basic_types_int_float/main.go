package main

import (
	"fmt"
)

func main(){
	
	views1 := 100
	views2 := 200
	totalViews := views1 + views2
	
	likes := 10.5
	likes++
	likes++
	avgViews := totalViews/2

	fmt.Println(totalViews, avgViews, likes)

}
