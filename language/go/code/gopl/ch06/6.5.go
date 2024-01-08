package ch06

import "fmt"

var BIT = 32 << (^uint(0) >> 63)

type IntSetWithBit struct {
	words []uint64
}

// Add adds the non-negative value x to the set.
func (s *IntSetWithBit) Add(x int) {
	word, bit := x/BIT, uint(x%BIT)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
		fmt.Println(len((s.words)))
	}
	s.words[word] |= 1 << bit
	fmt.Println(s.words)
}
