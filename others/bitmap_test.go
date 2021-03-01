package main

import (
	"math/rand"
	"sort"
	"testing"
)

const (
	n = 10_000_000
	k = 1_000_000
)

// 数据
var data3 [k]int

// 随机生成小于n的k个数
func init() {
	var total [n]int
	for i := range total {
		total[i] = i
	}
	for i := 0; i < k; i++ {
		ri := rand.Intn(n)
		total[i], total[ri] = total[ri], total[i]
	}
	copy(data3[:], total[:])
}

// 位图排序
func bitmapSort(s []int) {
	// 桶要装得下数据中可能出现的最大数
	var bucket [n]bool
	for i := range s {
		bucket[s[i]] = true
	}
	// 从桶中复制排序信息
	var cur int
	for i := range bucket {
		if bucket[i] {
			s[cur] = i
			cur++
		}
	}
}

func BenchmarkBitmapSort(b *testing.B) {
	s := make([]int, k)
	copy(s, data3[:])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitmapSort(s)
	}
}

func BenchmarkLibSort(b *testing.B) {
	s := make([]int, k)
	copy(s, data3[:])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}
