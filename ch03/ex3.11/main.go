// comma inserts commas in a non-negative decimal integer string. Write a non-recursive program to do this.
// Enhance comma so that it deals correclty with floating-pont numbers and an optional sign.
// e.g. "12345" -> "1,2,3,4,5"
package main

import "bytes"
import "fmt"

func comma(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteByte(byte(v))
		if i != (len(s) - 1) {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	str := "12345"
	fmt.Println(comma(str))
}
