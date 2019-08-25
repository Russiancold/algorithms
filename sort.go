package algorithms

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	a1 := MergeSort(a[:len(a)/2])
	a2 := MergeSort(a[len(a)/2:])

	b := make([]int, 0, len(a))

	j, k := 0, 0
	for i := 0; i < cap(b); i++ {
		if j == len(a1) {
			return append(b, a2[k:]...)
		}

		if k == len(a2) {
			return append(b, a1[j:]...)
		}

		if a1[j] <= a2[k] {
			b = append(b, a1[j])
			j++
		} else {
			b = append(b, a2[k])
			k++
		}
	}
	return b
}
