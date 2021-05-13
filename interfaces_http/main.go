package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type geometry interface {
//     area() float64
//     perim() float64
// }

// type rect struct{
// 	height, width float64
// }

// type circle struct{
// 	radius float64
// }

// func main(){
// 	r := rect{
// 		height: 10,
// 		width: 2.5,
// 	}
// 	c := circle{
// 		radius: 1.5,
// 	}

// 	measure(r)
// 	measure(c)
// }

// func (r rect) area() float64{
// 	return r.height*r.width
// }

// func (r rect) perim() float64{
// 	return 2*(r.height+r.width)
// }

// func (c circle) area() float64{
// 	return c.radius*c.radius
// }

// func (c circle) perim() float64{
// 	return 2*3.14*c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())

// }

type logWriter struct{}


func main() {

	resp, err := http.Get("http://google.com")
	if err!=nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(lw, resp.Body)


}

func (logWriter) Write(bs []byte) (n int, err error){
	fmt.Println(string(bs))
	fmt.Println("Write ", len(bs))
	return len(bs), nil
}