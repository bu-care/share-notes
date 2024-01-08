package ch07

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

// func (t *tree) String() string {
// 	var s string
// 	//当子树为nil时，返回空字符串
// 	if t == nil {
// 		return s
// 	}
// 	//每一个树的遍历顺序都是递归左中右
// 	s += t.left.String()
// 	s += " " + strconv.Itoa(t.value)
// 	s += t.right.String()

// 	return s
// }

// 实现Stringer接口
func (tree *tree) String() string {
	var values []int
	values = appendValues(values, tree)
	return fmt.Sprint(values)
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		//等价于遍历树，顺序都是递归左中右
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	fmt.Println("t.value", t.value)
	// 添加子树的时候都是比节点小的值放左边，比根节点大的放右边
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func addAll(t *tree, value []int) *tree {
	for _, v := range value {
		add(t, v)
	}
	return t
}
