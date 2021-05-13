package main

import "fmt"

type shape interface {
    getArea()  float64
}

type triangle struct{
	height, base float64
}

type square struct{
	sideLength float64
}

func main(){
	
	s := square{sideLength: 10}
	t := triangle{height: 5, base:2.5}

	printArea(s)
	printArea(t)
}

func (t triangle) getArea() float64{
	return 0.5* t.base* t.height
}

func (s square) getArea() float64{
	return s.sideLength* s.sideLength
}


func printArea(s shape) {
	
	fmt.Println(s.getArea())

}
