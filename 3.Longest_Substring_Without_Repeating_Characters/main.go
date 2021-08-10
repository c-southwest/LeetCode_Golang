package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	//var start, end int
	//length := 0
	//if len(s) == 0 {
	//	return 0
	//}
	//
	//m := map[string]int{}
	//
	//for k, v := range s {
	//	if index, ok := m[string(v)]; ok == false {
	//
	//		m[string(v)] = k
	//	} else {
	//		// index
	//		if length < end-start+1 {
	//			length = end - start + 1
	//		}
	//		for i := start; i <= index; i++ {
	//			delete(m, string(s[i]))
	//		}
	//		start = index + 1
	//		m[string(v)] = k
	//	}
	//	end = k
	//}
	//if length < end-start+1 {
	//	length = end - start + 1
	//}
	//return length

	var start, end int
	if len(s) == 0 {
		return 0
	}
	// 为0的情况排除，所以长度至少为1
	longest := 1
	m := map[string]int{}
	for k, v := range s {
		//如果没有键，或者值是老旧的
		if index, ok := m[string(v)]; ok == false || index < start {
			m[string(v)] = k
		} else {
			m[string(v)] = k //更新值
			//在更新start坐标之前，先把本轮长度给结算了
			if longest < end-start+1 {
				longest = end - start + 1
			}
			start = index + 1

		}
		end = k
	}
	//最后一轮循环后还没比较，需要再比较一次
	if longest < end-start+1 {
		longest = end - start + 1
	}
	return longest
}

func main() {
	str := "abcabcbb"
	r := lengthOfLongestSubstring(str)
	fmt.Println(r)
}
