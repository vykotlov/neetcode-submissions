func isValid(s string) bool {
	stack := RuneStack{}
	pairs := map[rune]rune{'{': '}', '[': ']', '(': ')'}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack.Push(r)
			continue
		}

		if bracket, ok := stack.Pop(); ok {
			if pairs[bracket] == r {
				continue
			}
		}

		return false
	}

	return stack.Len() == 0
}

type RuneStack struct {
	items []rune
}

func (s *RuneStack) Push(r rune) {
	s.items = append(s.items, r)
}

func (s *RuneStack) Pop() (rune, bool) {
	if s.Len() == 0 {
		return 0, false
	}

	n := len(s.items) - 1
	r := s.items[n]
	s.items = s.items[:n]

	return r, true
}

func (s *RuneStack) Len() int {
	return len(s.items)
}
