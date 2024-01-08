package ch06

import "fmt"

func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			// fmt.Printf("j: %d, word: %d, word&: %b\n", j, word, word&(1<<uint(j)))
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	fmt.Println("elems: ", elems)
	return
}
