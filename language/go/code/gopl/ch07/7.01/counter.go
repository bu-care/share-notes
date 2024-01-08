package counter

import (
	"bufio"
	"bytes"
	"fmt"
)

type Counter struct {
	LineCount int
	WordCount int
	ByteCount int
}

func (c *Counter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// fmt.Println(c)
		c.WordCount++
	}
	Lscanner := bufio.NewScanner(bytes.NewReader(p))
	Lscanner.Split(bufio.ScanLines)
	for Lscanner.Scan() {
		c.LineCount++
	}
	Bscanner := bufio.NewScanner(bytes.NewReader(p))
	Bscanner.Split(bufio.ScanBytes)
	for Bscanner.Scan() {
		c.ByteCount++
	}

	n := len(p)
	fmt.Println(n, string(p))

	return n, nil
}
