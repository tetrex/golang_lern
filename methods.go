package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt((*v).X*v.X + (*v).Y*v.Y)
}

func (v *Vertex) Abs2() float64 {
	return ((*v).X*2 + (*v).Y*2)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v2 := Vertex{2, 4}
	fmt.Println(v2.Abs2())
}
