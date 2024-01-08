package main

import (
	"fmt"
	"io"
)

type count struct {
	w io.Writer
	n int64
}

// Write 一个必须要实现的功能就是计算写入的byte数量，返回len(p), err
func (c *count) Write(p []byte) (int, error) {
	num, err := c.w.Write(p)
	c.n += int64(num)
	return num, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &count{
		w: w,
	}
	return c, &c.n
}

func main() {
	// w := io.Discard
	// 不用 io.Discard，也可以自己写一个实现 io.Writer 的结构体
	var w io.Writer = myDiscard{}
	c, n := CountingWriter(w)
	fmt.Println("before write: ", c, *n)
	num, _ := fmt.Fprintf(c, "counting xbu")
	fmt.Println(num, c, *n)
}

type myDiscard struct{}

func (myDiscard) Write(p []byte) (int, error) {
	return len(p), nil
}
