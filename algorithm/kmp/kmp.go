package kmp

func Index(text, sub string) int {
	next := getNext(sub)
	i, j := 0, 0
	for i < len(text) {
		if sub[j] == text[i] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = next[j-1]
			}
		}

		if j == len(sub) {
			return i - j
		}
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	i, j := 0, 1
	for j < len(s) {
		if s[i] == s[j] {
			i++
			next[j] = i
			j++
		} else {
			if i == 0 {
				j++
			} else {
				i = next[i-1]
			}
		}
	}
	return next
}
