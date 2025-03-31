package main

import (
	"fmt"
	"strings"
	"test/hr"
)

func main() {
	rspBytes, err := hr.RequestDw("http://idc.ntsgw.woa.com/ntsgw/api/esb/dos-interface-server/open-api/config/hrdw/dw-api-private-er-iegcros-org-base/IEG-CROS/file")
	rspStr := strings.Trim(string(rspBytes), "\n")
	fmt.Println("Converted time:", err)
	fmt.Println("Converted time:", rspStr)
}

// Diff 求差集 source-并集
func Diff(source, diff []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := SliceIntersect(source, diff)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range source {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// SliceIntersect 求交集
func SliceIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}
