package main

import (
	"algorithm/tool"
)

func init() {
	args := []int{3, 1, 2, 4, 5}
	tool.RunAndPick(&args, quick, args, 0, len(args))
}

// 左闭右开
func quick(s []int, lo, hi int) {
	if lo < hi-1 {
		p := partition(s, lo, hi)
		quick(s, lo, p)
		quick(s, p+1, hi)
	}
}

func partition(nums []int, lo, hi int) int {
	p := lo
	pv := nums[p]
	for {
		for {
			lo++
			if lo >= hi || // hi及其右侧元素确定都大于pv，所以lo没必要继续往右走了
				nums[lo] > pv { // 找到了某个不符合lo界限要求的元素
				break
			}
		}

		for {
			hi--
			if hi == p || // hi的左边界，这里hi一定要保证它及其右侧元素都大于pv，不能以lo为界
				nums[hi] < pv { // 找到了不符合hi界限要求的元素
				break
			}
		}

		if lo >= hi { // 判断是否该终止
			break
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}

	// 到这里时，lo左侧元素都不大于pv，hi上及其右侧元素都不小于pv
	// lo所处位置元素可能大于pv，这时交换至左侧是不符合要求的
	// 而hi的终止条件确保要么hi==p要么hi元素小于pv，可换至左侧
	nums[p], nums[hi] = nums[hi], nums[p]
	return hi
}
