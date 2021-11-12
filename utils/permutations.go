package utils

// Permutations of n integers, to be used as index for slice permutations.
// Uses Heap's Algorithm (thanks wikipedia)
func Permutations(n int) <-chan []int {
	ch := make(chan []int)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i
	}

	go func() {
		intPermutationsRecursor(n, a, ch)
		close(ch)
	}()

	return ch
}

func intPermutationsRecursor(k int, a []int, ch chan<- []int) {
	if k == 1 {
		output := make([]int, len(a))
		copy(output, a)
		ch <- output
	} else {
		intPermutationsRecursor(k-1, a, ch)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				a[i], a[k-1] = a[k-1], a[i]
			} else {
				a[0], a[k-1] = a[k-1], a[0]
			}
			intPermutationsRecursor(k-1, a, ch)
		}
	}
}

// Creates a channel that returns combinations of the first n integers
// in groups of length m. To be used as indices for combinations of slices.
func Combinations(n int, m int) <-chan []int {
	ch := make(chan []int)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i
	}

	go func() {
		prefix := []int{}
		combinationsRecursor(prefix, m, a, ch)
		close(ch)
	}()

	return ch
}

func combinationsRecursor(prefix []int, m int, a []int, ch chan<- []int) {
	if m == 1 {
		for _, c := range a {
			l := len(prefix)
			result := make([]int, l+1)
			copy(result, prefix)
			result[l] = c
			ch <- result
		}
		return
	}

	if m == len(a) {
		result := make([]int, len(prefix)+m)
		copy(result, prefix)
		copy(result[len(prefix):], a)
		ch <- result
		return
	}

	l := len(prefix)
	newPrefix := make([]int, l+1)

	copy(newPrefix, prefix)
	newPrefix[l] = a[0]

	// combinations with this element
	combinationsRecursor(newPrefix, m-1, a[1:], ch)

	// combinations without this element
	combinationsRecursor(prefix, m, a[1:], ch)
}
