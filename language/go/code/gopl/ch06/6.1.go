package ch06

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
		fmt.Println(len((s.words)))
	}
	// 1、1左移了bit位，相当于乘以2的bit次方，1 << bit就等于 1*2**9=512
	// 2、在商为索引的位置上只会存储一个值，如112、113两个数都只会存在索引1的位置上，
	// 3、后存储的与当前值进行或操作后把当前值覆盖掉，以64位二进制来存储，如此就可以保证一个位置的数值就可以存储64个数（如索引1位置存储64-127）
	s.words[word] |= 1 << bit
	fmt.Println(s.words)
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, e := range s.Elems() {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", e)
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int { // return the number of elements
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			// 判断word的二进制第j位是否为1，如果为1, 则代表这个位置存在数字
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) { // remove x from the set
	word, bit := x/64, uint(x%64)

	s.words[word] ^= 1 << bit
	// 判断word是否越界
	if word < len(s.words) {
		// 如果没有越界, 则将word的第bit位置为0
		// 先进行与操作，再进行异或操作（不同为1，相同为0），将代表 x 的值从64位二进制位置上去掉
		// 如果集合中存在x，与操作不会变化，如果不存在x，就相当于先添加再删除
		s.words[word] ^= s.words[word] & (1 << bit)
		// <=> s.words[word] &= ^(1 << bit)，^(1 << bit)表示与0进行异或操作
	}
	fmt.Println("after remove: ", s.words)
}

func (s *IntSet) Clear() { // remove all elements from the set
	// for i := range s.words {
	// 	s.words[i] = 0
	// }
	s.words = nil
	fmt.Println("after Clear: ", s.words)
}

func (s *IntSet) Copy() *IntSet { // return a copy of the set
	var set IntSet
	set.words = append(set.words, s.words...)
	return &set
}
