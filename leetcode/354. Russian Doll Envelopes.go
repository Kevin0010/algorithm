package main

import "algorithm/tool"

func init() {
	tool.Run(maxEnvelopes, [][]int{{5, 8}, {6, 4}, {6, 7}, {2, 3}})
}

func maxEnvelopes(envelopes [][]int) int {
	l := len(envelopes)
	quickSortByWidth(envelopes, 0, l)
	return 0
}

func quickSortByWidth(envelops [][]int, lo, hi int) {
	if lo < hi-1 {
		p := partitionByWidth(envelops, lo, hi)
		quickSortByWidth(envelops, 0, p)
		quickSortByWidth(envelops, p+1, hi)
	}
}

func partitionByWidth(envelops [][]int, lo, hi int) int {
	p := lo
	pw := envelops[p][0]
	ph := envelops[p][1]
	lo++
	hi--
	for lo < hi {
		for {
			if envelops[lo][0] > pw || envelops[lo][0] == pw && envelops[lo][1] >= ph || lo < hi {
				break
			}
			lo++
		}
		for envelops[hi][0] > pw {
			hi--
		}
		if lo >= hi {
			break
		}
		envelops[lo], envelops[hi] = envelops[hi], envelops[lo]
	}
	envelops[p], envelops[hi] = envelops[hi], envelops[p]
	return hi
}
