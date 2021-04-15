//Interfaces Challenge
//io.Writer interface
//Write a struct called Capper that has a field to another io.Writer
//This transforms everything written to uppercase.package interchall
//Capper should implement io.Writer
package main

import (
	"fmt"
	"io"
	"os"
)

// struct, Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
