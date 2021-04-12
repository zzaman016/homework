package main

import "fmt"

func main(){
	var n, y, p, m, s float64
	fmt.Scan(&n, &y, &p)
	m = (n*p*(1+p)*y)/(12*((1+p)*y-1))
	fmt.Println("Месячная выплата равна: ", m)
	s=(m*12)*y
	fmt.Println("Суммарная выплата: ", s)
}
