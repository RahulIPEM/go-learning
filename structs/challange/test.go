package main

import (
	"fmt"
	"log"
)

// Point is a 2D point
type Point struct {
	x int
	y int
}

// Move moves the point
func (p *Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}

// Square square
type Square struct {
	Center Point // center of square
	Length int   // length
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

// Area return the square's area
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
