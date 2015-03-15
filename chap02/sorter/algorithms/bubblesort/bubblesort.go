package bubblesort

func Sort(vals []int) {

	for i := len(vals) - 1; i > 0; i-- {
		jump := true

		for j := 0; j < i; j++ {
			if vals[j] > vals[j+1] {
				vals[j], vals[j+1] = vals[j+1], vals[j]
				jump = false
			}
		}

		if jump {
			break
		}
	}
}
