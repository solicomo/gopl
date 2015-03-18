package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopl/chap02/sorter/algorithms/bubblesort"
	"gopl/chap02/sorter/algorithms/qsort"
	"io"
	"os"
	"strconv"
)

type LineTooLongError struct {
	lnum int
	line string
}

func (e *LineTooLongError) Error() string {
	return "Line " + strconv.Itoa(e.lnum) + " is too long."
}

func (e *LineTooLongError) Line() string {
	return e.line
}

func readValues(inf *string) (vals []int, err error) {
	ifile, err := os.Open(*inf)

	if err != nil {
		return
	}

	defer ifile.Close()

	rd := bufio.NewReader(ifile)
	vals = make([]int, 0)

	for i := 0; ; i++ {
		line, isPre, er := rd.ReadLine()

		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}

		str := string(line)

		if isPre {
			err = &LineTooLongError{i, str}
			break
		}

		val, er := strconv.Atoi(str)

		if er != nil {
			err = er
			break
		}

		vals = append(vals, val)
	}

	return
}

func writeValues(vals []int, ouf *string) (err error) {
	ofile, err := os.Create(*ouf)

	if err != nil {
		return
	}

	defer ofile.Close()

	for _, val := range vals {
		str := strconv.Itoa(val)
		ofile.WriteString(str + "\n")
	}

	return
}

func main() {
	inf := flag.String("i", "in", "Input File")
	ouf := flag.String("o", "out", "Output File")
	alg := flag.String("a", "qsort", "Sort Algorithm")

	flag.Parse()

	fmt.Println(*inf, "=", *alg, "=>", *ouf)

	vals, err := readValues(inf)

	if err != nil {
		fmt.Println("read file", *inf, "failed:", err)
		return
	}

	switch *alg {
	case "qsort":
		qsort.Sort(vals)
	case "bubblesort":
		bubblesort.Sort(vals)
	default:
		fmt.Println("Sort algorithm", *alg, "is invalid. Only qsort and bubblesort are supported.")
		return
	}

	err = writeValues(vals, ouf)

	if err != nil {
		fmt.Println("write file", *ouf, "failed:", err)
	}

}
