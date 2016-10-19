package strings2

import (
	"bytes"
	"strconv"
)

// JoinInt just like Join on standard lib of strings, this func is used to int slice
// and convert int slice `ints` to string seperated by `sep`
func JoinInt(ints []int, sep string) string {
	var buffer bytes.Buffer
	for i, v := range ints {
		if i != 0 {
			buffer.WriteString(sep)

		}
		buffer.WriteString(strconv.Itoa(v))

	}
	return buffer.String()
}
