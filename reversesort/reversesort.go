package reversesort

// ReverseSort sorts a list of distinct integers in increasing order
// will return the cost of the sort operation
func ReverseSort(l *[]int) (cost int) {
	for i := 0; i < len(*l)-1; i++ {
		j := i + minIndex((*l)[i:])
		cost += j - i + 1
		reverse(l, i, j)
	}
	return
}

// minIndex returns the index of the minimum value in the specified range
// will return 0 if an empty or nil slice is passed as the argument
func minIndex(l []int) (i int) {
	for _i, v := range l {
		if v < l[i] {
			i = _i
		}
	}
	return
}

// reverse will reverse the slice in the specified range
func reverse(l *[]int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}
