package intheap

type BadHeapError string

func (e BadHeapError) Error() string {
	return string(e)
}

const ErrBadPushType BadHeapError = "IntHeap.Push expects int"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) error {
	value, ok := x.(int)
	if !ok {
		return ErrBadPushType
	}

	*h = append(*h, value)

	return nil
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func (h *IntHeap) Sort() {
	length := len(*h)
	for i := range length - 1 {
		for j := range length - 1 - i {
			if (*h)[j] > (*h)[j+1] {
				h.Swap(j, j+1)
			}
		}
	}
}
