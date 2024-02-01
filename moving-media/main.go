// https://pdfcoffee.com/problem-solving-2-3-pdf-free.html
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Have the function MovingMedian(arr) read the array of numbers stored in arr
// which will contain a sliding window size, N, as the first element in the array
// and the rest will be a list of numbers. Your program should return the Moving
// Median for each element based on the element and its N-1 predecessors, where N
// is the sliding window size. The final output should be a string with the moving
// median corresponding to each entry in the original array separated by commas.

// Note that for the first few elements(until the window size is reached), the median
// is computed on a smaller number of entries.For example : if arr is [3, 1, 3, 5, 10,
// 6, 4, 3, 1] then your program should output "1,2,3,5,6,6,4,3".

func MovingMedian(arr []int) string {
	windowSize := arr[0]
	output := &strings.Builder{}

	// Note that for the first few elements(until the window size is reached), the median
	// is computed on a smaller number of entries.
	for ix := 1; ix <= windowSize; ix++ {
		if ix+1 > len(arr) {
			break
		}
		output.WriteString(strconv.Itoa(calculateMedian(arr[1 : ix+1])))
		output.WriteString(",")
	}

	// Calculation of the other moving medias
	for ix := windowSize - 1; ix <= len(arr)-windowSize; ix++ {
		output.WriteString(strconv.Itoa(calculateMedian(arr[ix : ix+windowSize])))
		output.WriteString(",")
	}

	ans := output.String()
	return ans[:len(ans)-1]
}

func calculateMedian(arr []int) int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	sort.Ints(copyArr)
	if len(copyArr)%2 == 0 {
		return (copyArr[len(copyArr)/2-1] + copyArr[len(copyArr)/2]) / 2
	}
	return copyArr[len(copyArr)/2]
}

func main() {
	arr := []int{3, 1, 3, 5, 10, 6, 4, 3, 1}
	fmt.Println("Input:", arr)
	fmt.Println("Output:", MovingMedian(arr))
	fmt.Println()
	arr = []int{5, 2, 4, 6}
	fmt.Println("Input:", arr)
	fmt.Println("Output:", MovingMedian(arr))
	fmt.Println()
	arr = []int{3, 0, 0, -2, 0, 2, 0, -2}
	fmt.Println("Input:", arr)
	fmt.Println("Output:", MovingMedian(arr))
}
