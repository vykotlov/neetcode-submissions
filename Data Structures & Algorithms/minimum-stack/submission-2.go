type MinStack struct {
	items []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	newMin := val

	if len(s.items) > 0 {
		currentMin := s.mins[len(s.mins)-1]
		if currentMin < newMin {
			newMin = currentMin
		}
	}

	s.items = append(s.items, val)
	s.mins = append(s.mins, newMin)
}

func (s *MinStack) Pop() {
	n := len(s.items) - 1

	s.items = s.items[:n]
	s.mins = s.mins[:n]
}

func (s *MinStack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}
