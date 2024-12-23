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
	"slices"
	"testing"

	"github.com/the-flx/flx.go"
)

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
	if !reflect.DeepEqual(result, []int{2, 3, 4}) {
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

func TestSlice1(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	result := slices.Delete(slice, 0, 2)
	if !reflect.DeepEqual(result, []int{3, 4}) {
		t.Errorf("Test `Slice` 1: %v", result)
	}
}

func TestGetHeatmapStr1(t *testing.T) {
	result := flx.GetHeatmapStr("switch-to-buffer", nil)
	if !reflect.DeepEqual(result, []int{82, -4, -5, -6, -7, -8, -9, 79, -7, -8, 76, -10, -11, -12, -13, -13}) {
		t.Errorf("Test `TestGetHeatmapStr` 1: %v", result)
	}
}

func TestBiggerSublist1(t *testing.T) {
	result := flx.BiggerSublist([]int{1, 2, 3, 4}, nil)
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4}) {
		t.Errorf("Test `BiggerSublit` 1: %v", result)
	}
}

func TestBiggerSublist2(t *testing.T) {
	val := 2
	result := flx.BiggerSublist([]int{1, 2, 3, 4}, &val)
	if !reflect.DeepEqual(result, []int{3, 4}) {
		t.Errorf("Test `BiggerSublit` 2: %v", result)
	}
}

func TestScore1(t *testing.T) {
	result := flx.Score("switch-to-buffer", "stb")
	if !reflect.DeepEqual(result.Indices, []int{0, 7, 10}) ||
		result.Score != 237 ||
		result.Tail != 0 {
		t.Errorf("Test `Score` 1: %v", result)
	}
}

func TestScore2(t *testing.T) {
	result := flx.Score("TestSomeFunctionExterme", "met")
	if !reflect.DeepEqual(result.Indices, []int{6, 16, 18}) ||
		result.Score != 57 ||
		result.Tail != 0 {
		t.Errorf("Test `Score` 2: %v", result)
	}
}

func TestScore3(t *testing.T) {
	result := flx.Score("MetaX_Version", "met")
	if !reflect.DeepEqual(result.Indices, []int{0, 1, 2}) ||
		result.Score != 211 ||
		result.Tail != 2 {
		t.Errorf("Test `Score` 3: %v", result)
	}
}
