package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius* c.radius
}

func (r Rectangle) area() float64{
	return r.width * r.height
}

func getArea(s Shape) float64{
	return s.area()
}

func main(){
	circle := Circle{x:0,y:0,radius:34}
	rectangle := Rectangle{width:340, height:450}

	fmt.Printf("Area of Circle: %f\n",getArea(circle))
	fmt.Printf("Area of Rectangle: %f\n",getArea(rectangle))
}