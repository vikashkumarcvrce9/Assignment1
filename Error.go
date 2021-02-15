package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(value float64) (float64,error)
if(value<0){
	return 0,errors.new("Negative nuumber passed to Sqrt")
}
return math.Sqrt(value),null
}
func main(){
	result,err=sqrt(64)
	if err!=null {
		fmt.Println(value)
		
	}else{
		fmt.Println(result)
	}
}