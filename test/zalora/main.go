package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValue(k, n int, w, v []int) int {
	if n == 0 {
		return 0
	}
	F := make([]int, k+1)

	for i := 1; i < n; i++ {
		for j := k; j >= w[i]; j-- {
			F[j] = max(F[j], F[j-w[i]]+v[i])
		}
	}
	fmt.Println(F)
	return F[k]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	knTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	kTemp, err := strconv.ParseInt(knTemp[0], 10, 64)
	checkError(err)
	nTemp, err := strconv.ParseInt(knTemp[1], 10, 64)
	checkError(err)
	k := int(kTemp)
	n := int(nTemp)

	wTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	vTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var w, v []int
	w = append(w, 0)
	v = append(v, 0)

	for i := 0; i < int(n); i++ {
		wItemTemp, err := strconv.ParseInt(wTemp[i], 10, 64)
		checkError(err)
		vItemTemp, err := strconv.ParseInt(vTemp[i], 10, 64)
		checkError(err)
		w = append(w, int(wItemTemp))
		v = append(v, int(vItemTemp))
	}

	fmt.Println(maxValue(k, n, w, v))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
