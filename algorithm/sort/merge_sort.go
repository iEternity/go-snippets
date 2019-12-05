package sort

func MergeSort(data []int) {
	if len(data) <= 1 {
		return
	}

	mid := (len(data) - 1) / 2
	MergeSort(data[:mid+1])
	MergeSort(data[mid+1:])

	ans := make([]int, 0, len(data))
	i, j := 0, mid+1
	for i <= mid && j < len(data) {
		if data[i] < data[j] {
			ans = append(ans, data[i])
			i++
		} else {
			ans = append(ans, data[j])
			j++
		}
	}

	for i <= mid {
		ans = append(ans, data[i])
		i++
	}

	for j < len(data) {
		ans = append(ans, data[j])
		j++
	}

	copy(data, ans)
}
