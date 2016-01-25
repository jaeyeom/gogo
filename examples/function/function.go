package function

import (
	"bufio"
	"fmt"
	"io"
)

// AddOne increases each element value in nums by 1.
func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

// WriteTo writes each element in lines to w.
func WriteTo(w io.Writer, lines []string) (n int64, err error) {
	for _, line := range lines {
		var nw int
		nw, err = fmt.Fprintln(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return
}

// ReadFrom calls f for each line from r.
func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
