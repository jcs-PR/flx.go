/**
 * $File: flx_test.go $
 * $Date: 2024-12-22 09:38:50 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package flx_test

import (
	"reflect"
	"testing"

	"github.com/the-flx/flx.go"
)

/* Test helper */

func compareList(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

/* Test */

func TestWord1(t *testing.T) {
	ch := 'c'
	result := flx.Word(&ch)
	if result == true {
		return
	}
	t.Errorf("Test `Word` 1: %v", result)
}

func TestWord2(t *testing.T) {
	ch := ' '
	result := flx.Word(&ch)
	if result == false {
		return
	}
	t.Errorf("Test `Word` 2: %v", result)
}

func TestIncVec1(t *testing.T) {
	vec := []int{1, 2, 3}
	inc := 1
	beg := 0
	end := 3
	result := flx.IncVec(vec, &inc, &beg, &end)
	if !compareList(result, []int{2, 3, 4}) {
		t.Errorf("Test `IncVec` 2: %v", result)
	}
}

func TestGetHashMapForString1(t *testing.T) {
	result := flx.GetHashForString("switch-to-buffer")
	if !reflect.DeepEqual(result, map[int][]int{
		114: {15},
		101: {14},
		102: {12, 13},
		117: {11},
		98:  {10},
		45:  {6, 9},
		111: {8},
		116: {3, 7},
		104: {5},
		99:  {4},
		105: {2},
		119: {1},
		115: {0},
	}) {
		t.Errorf("Test `GetHashForString` 1: %v", result)
	}
}
