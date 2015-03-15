package qsort

func Sort(vals []int) {
	qsort(vals, 0, len(vals)-1)
	return
}

func qsort(vals []int, l, r int) {
	if l >= r {
		return
	}

	val := vals[l]
	i, j:= l, r

	for i < j {
		for j > i && vals[j] >= val {
			j--
		}

		vals[i] = vals[j]
		
		for i < j && vals[i] <= val {
			i++
		}

		vals[j] = vals[i]
	}

	vals[i] = val

	qsort(vals, l, i)
	qsort(vals, i+1, r)
}

