package main

import "strings"

func init() {
	run(convert, "AB", 1)
}

func convert(s string, numRows int) string {
	l := len(s)
	roundLen := numRows + numRows - 2
	if roundLen == 0 {
		roundLen = 1
	}
	var sb strings.Builder
	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		secInRow := rowIdx > 0 && rowIdx < numRows-1
		for idx, col := rowIdx, roundLen; idx < l; idx += roundLen {
			sb.WriteByte(s[idx])
			if secInRow {
				ri := col - idx + (col - roundLen)
				if ri < l {
					sb.WriteByte(s[ri])
				}
				col += roundLen
			}
		}
	}
	return sb.String()
}
