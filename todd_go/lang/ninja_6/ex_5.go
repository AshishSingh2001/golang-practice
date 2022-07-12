package main

// import (
// 	"fmt"
// 	"math"
// )

// type square struct {
// 	side float64
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	r := c.radius
// 	return math.Pi * r * r
// }

// func (s square) area() float64 {
// 	r := s.side
// 	return r * r
// }

// type shape interface {
// 	area() float64
// }

// func info(s shape) float64 {
// 	return s.area()
// }

// func main() {
// 	s := square{
// 		side: 3,
// 	}
// 	c := circle{
// 		radius: 1,
// 	}
// 	fmt.Println(info(s), info(c))
// }
