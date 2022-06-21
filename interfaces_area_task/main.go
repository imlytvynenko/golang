package main

import "fmt"


type shape interface{
	getArea() float64
}

type square struct{
	sideLenght float64
}
type triangle struct{
	base float64
	height float64
}

func main() {
	sq := square{sideLenght: 5.05}
	tr := triangle{base: 3.78, height: 4.05}
	pritnArea(sq)
	pritnArea(tr)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func pritnArea(s shape) {
	fmt.Println("Square is - ", s.getArea())
}

