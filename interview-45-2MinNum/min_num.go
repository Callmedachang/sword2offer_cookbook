package interview_45_2MinNum

import (
	"fmt"
	"strings"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
*/

func MinNumber(nums []int) string {
	Sort(nums, 0, len(nums)-1)
	res := strings.Builder{}
	for i := range nums {
		res.WriteString(fmt.Sprint(nums[i]))
	}
	return res.String()
}

func Sort(nums []int, begin, end int) {
	if begin < end {
		index := partition(nums, begin, end)
		Sort(nums, begin, index-1)
		Sort(nums, index+1, end)
	}
}

func partition(nums []int, begin, end int) int {
	pivot, index := nums[end], begin
	for i := begin; i < end; i++ {
		if Bigger(nums[i], pivot) {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[index], nums[end] = nums[end], nums[index]
	return index
}

func Bigger(n1, n2 int) bool {
	if fmt.Sprintf("%d%d", n1, n2) > fmt.Sprintf("%d%d", n2, n1) {
		return false
	}
	return true
}
