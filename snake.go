package snake

import "math"

type SnakeSegment struct {
	Point
}

type Snake struct {
	head    *Point
	heading float64
	speed   float64
	body    [1000]Point
	size    float64
}

func (s *Snake) Move() {
	// First, move the snake's head
	dx := math.Sin(s.heading) * 3
	dy := math.Cos(s.heading) * 3

	s.body[0].X += dx
	s.body[0].Y += dy
	for i := 1; i < s.Length(); i++ {
		s.body[i].X += (s.body[i-1].X - s.body[i].X) * 0.4
		s.body[i].Y += (s.body[i-1].Y - s.body[i].Y) * 0.4
	}
}

func (s *Snake) Length() int {
	return 15 + int(s.size)
}

func (s *Snake) Grow() {
	oldLength := s.Length()
	s.size += 1
	if newLength := s.Length(); newLength > oldLength {
		s.body[newLength-1] = s.body[oldLength-1]
	}
}

func (s *Snake) TurnLeft() {
	s.heading += 0.05
}
func (s *Snake) TurnRight() {
	s.heading -= 0.05
}
