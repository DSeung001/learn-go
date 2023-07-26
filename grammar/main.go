package grammar

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// React struct, Shape interface implement 구현
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle struct, Shape interface implement 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 어떤 타입이든 인터페이스를 구현한 경우 인터페이스로 받을 수 있다
func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func main() {
	// 인터페이스 매개변수로 사용
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)

	// 인터페이스 타입 = interface type = 빈 인터페이스
	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)
}

func printIt(v interface{}) {
	fmt.Println(v) // Tom 출력
}
