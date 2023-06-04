package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	braketMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := braketMap[r]; ok {
			stack = append(stack, r)
			// 末尾の要素がペアでなければfalse
		} else if len(stack) == 0 || braketMap[stack[len(stack)-1]] != r {
			return false
		} else {
			// 末尾の要素がペアであれば削除
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// @lc code=end
