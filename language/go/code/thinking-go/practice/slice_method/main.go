package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	//两点间距离
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// 方法是类型的方法，不只是只有结构体才有方法，定义了一个 Path 的切片类型，也可以定义 Path 的方法
func (p Path) Distance() float64 {
	sum := 0.0
	for i, j := range p {
		if i > 0 {
			sum += p[i-1].Distance(j)
		}
	}
	return sum
}

func main() {
	//切片中赋初值时没有指明是 Point 也可以？
	points := Path{
		Point{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	////计算三角形周长
	fmt.Println(points.Distance())
}
