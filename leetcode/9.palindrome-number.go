package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {

	s := strconv.Itoa(x)

	l := len(s)

	for i := 0;i<=l/2;i++ {
		v1 := string(s[i]) 
		v2 := string(s[l - i - 1])

		if v1 != v2 {
			return false
		}
	}

	return true
}
// @lc code=end

