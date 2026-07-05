type MinStack struct {
	items []int
	min   int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)

	if val < this.min || len(this.items) == 1 {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	n := len(this.items) - 1

	val := this.items[n]
	this.items = this.items[:n]

	if len(this.items) == 0 {
		this.min = 0
		return
	}

	if val == this.min {
		newMin := this.items[0]

		for i := 1; i < len(this.items); i++ {
			if this.items[i] < newMin {
				newMin = this.items[i]
			}
		}

		this.min = newMin
	}
}

func (this *MinStack) Top() int {
	if len(this.items) == 0 {
		return 0
	}

	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
