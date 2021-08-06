package snake

type SnakeSegment struct {
	Point
}

type Snake struct {
	head *Point
	dx   float64
	dy   float64
	body [1000]Point
	size float64
	// vertical and horizontal momentum
	vmom, hmom float64
}

func (s *Snake) Move() {
	// First, move the snake's head
	s.head.X += s.hmom
	s.head.Y += s.vmom

	for i := 1; i < s.length(); i++ {
		s.body[i].X += (s.body[i-1].X - s.body[i].X) * (2 / s.size)
		s.body[i].Y += (s.body[i-1].Y - s.body[i].Y) * (2 / s.size)
	}
}

func (s *Snake) length() int {
	return int(s.size)
}

func (s *Snake) Grow() {
	oldLength := s.length()
	s.size += 1
	if newLength := s.length(); newLength > oldLength {
		s.body[newLength-1] = s.body[oldLength-1]
	}
}
