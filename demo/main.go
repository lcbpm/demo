package main

import (
	"fmt"
)

func main() {
	slice()
	chanOperator()
}

func chanOperator() {
	的微风微风微风微风
	number := make(chan bool)
	letter := make(chan bool)
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j >= 'Z' {
					done <- true
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					number <- true
				}
			}
		}
	}()

	number <- true

	for {
		select {
		case <-done:
			return
		}
	}
}

func slice() {
	// 声明和初始化
	var s1 []int
	s1 = append(s1, 1)
	s2 := make([]int, 5)
	s3 := []int{1, 2, 3}

	arr := [5]int{1, 2, 3, 4, 5}
	s4 := arr[:]

	// 复制
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src)

	// 追加
	s5 := []int{1, 2, 3}
	s5 = append(s5, 4, 5)

	// 多维切片
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var s6 [4]int

	fmt.Println(s1, s2, s3, arr, s4, src, dest, s5, matrix)
	fmt.Println(s1 == nil, s2 == nil, s3 == nil, s4 == nil, s5 == nil, s6)
	fmt.Println(s1 == nil, s2 == nil, s3 == nil, s4 == nil, s5 == nil)

}
