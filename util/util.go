package util

import "strconv"

// CovertPageToOffset 把页码和页数转换为开始和偏移量
// 2, 5 => 5, 5
func CovertPageToOffset(page string, size string) (int, int) {
	isize, _ := strconv.Atoi(size)
	ipage, _ := strconv.Atoi(page)
	return isize * (ipage - 1), isize
}
