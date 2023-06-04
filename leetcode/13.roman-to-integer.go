package main

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	vmap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	for _, r := range s {
		current := vmap[r]

		if current > prev {
			result += current - 2*prev
		} else {
			result += current
		}

		prev = current
	}

	return result
}

// @lc code=end
