/**
 * $File: util.go $
 * $Date: 2024-12-22 09:37:57 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package flx

import (
	"strings"
)

// Return true if `slice` contains the `item`.
func contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func toLower(ch rune) rune {
	return toRune(strings.ToLower(string(ch)))
}

func toRune(str string) rune {
	var runes = []rune(str)
	return runes[0]
}

func fillSlice(len int, defaultVal int) []int {
	var arr []int = []int{}
	for range len {
		arr = append(arr, defaultVal)
	}
	return arr
}

/* Dictionary */

func dictSet(result map[int][]Result, key *int, val []Result) map[int][]Result {
	if key == nil {
		return result
	}

	result[*key] = val

	return result
}

func dictGet[T any](dict map[int][]T, key *int) []T {
	if key == nil {
		return nil
	}

	val, ok := dict[*key]

	if ok {
		return val
	}

	return nil
}

func dictInsert(result map[int][]int, key int, val int) map[int][]int {
	var lst []int = result[key]
	result[key] = append([]int{val}, lst...)
	return result
}
