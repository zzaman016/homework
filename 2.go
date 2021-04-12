package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c, P,S, p float64
	fmt.Scan(&a,&b,&c)
	P=a+b+c
	fmt.Println("Периметр треугольника: ",P)
	p=P/2
	S=math.Sqrt((p*(p-a)*(p-b)*(p-c)))
	fmt.Print("Площадь треугольника: ", S)

}
