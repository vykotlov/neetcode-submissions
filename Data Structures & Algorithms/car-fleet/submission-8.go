func carFleet(target int, positions []int, speeds []int) int {
	speedsByStartPosition := make(map[int]int)
	for i := 0; i < len(positions); i++ {
		speedsByStartPosition[positions[i]] = speeds[i]
	}

	sort.Ints(positions)
	stack := NewStack(len(positions))
	for i := 0; i < len(positions); i++ {
		stack.Push(positions[i])
	}

	fleetsCount := 0
	lastArrivalTime := 0.0

	for stack.Length() > 0 {
		position := stack.Pop()
		speed := speedsByStartPosition[position]

		arrivalTime := float64(target-position) / float64(speed)

		if arrivalTime > lastArrivalTime {
			fleetsCount++
			lastArrivalTime = arrivalTime
		}
	}

	return fleetsCount
}

type Stack struct {
	items []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		items: make([]int, 0, capacity),
	}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

func (s *Stack) Length() int {
	return len(s.items)
}
