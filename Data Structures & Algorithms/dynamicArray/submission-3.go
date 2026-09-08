
type DynamicArray struct {
	backingArray []int
	l, c         int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		backingArray: make([]int, capacity),
		l:            0,
		c:            capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.backingArray[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.backingArray[i] = n
}


func (da *DynamicArray) Pushback(n int) {
	if da.c <= da.l {
		da.resize()
	}
		// we do simple
		da.backingArray[da.l] = n
		da.l++
}

func (da *DynamicArray) Popback() int {
	last := da.backingArray[da.l-1]
	da.l -= 1
	return last

}

func (da *DynamicArray) resize() {
	init := len(da.backingArray)
	newBa := make([]int, 2*init)
	da.l = da.c
	da.c = 2 * init

	for i, el := range da.backingArray {
		newBa[i] = el
	}

	da.backingArray = newBa
}

func (da *DynamicArray) GetSize() int {
	return da.l
}

func (da *DynamicArray) GetCapacity() int {
	return da.c
}
