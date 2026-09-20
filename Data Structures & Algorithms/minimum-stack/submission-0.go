type MinStack struct {
	bArr [][2]int
}

func Constructor() MinStack {
	return MinStack{
		bArr: [][2]int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.bArr) == 0 {
		this.bArr = append(this.bArr, [2]int{val, val})
	} else {
		this.bArr = append(this.bArr, [2]int{val, min(val, this.bArr[len(this.bArr)-1][1])})
	}
}

func (this *MinStack) Pop() {
	this.bArr = this.bArr[:len(this.bArr)-1]
}

func (this *MinStack) Top() int {
	return this.bArr[len(this.bArr)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.bArr[len(this.bArr)-1][1]
}

