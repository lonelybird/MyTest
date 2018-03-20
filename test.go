package main

import "fmt"

func main(){

	const(
		a = iota
		b
		c 
		d = "ha"
		e
		f = 100
		g 
		m 
		n
		h = iota
		i 
	)
	
	fmt.Printf(" a = %d\n", a)
	fmt.Println(a,b,c,d,e,f,g,m,n,h,i)
}