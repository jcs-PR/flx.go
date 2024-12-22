/**
 * $File: flx.go $
 * $Date: 2024-12-22 07:52:08 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package flx

import (
	"strings"
)

var wordSeparators = []rune{' ', '-', '_', ':', '.', '/', '\\'}

var defaultScore = -35

// Check if `ch` is a word character.
func Word(ch *rune) bool {
	if ch == nil {
		return false
	}
	return !contains(wordSeparators, *ch)
}

// Check if `ch` is an uppercase character.
func Capital(ch *rune) bool {
	var s = string(*ch)
	return Word(ch) && s == strings.ToUpper(s)
}

// Check if `lastCh` is the end of a word and `ch` the start of the next.
//
// This function is camel-case aware.
func Boundary(lastCh *rune, ch *rune) bool {
	if lastCh == nil {
		return false
	}

	if !Capital(lastCh) && Capital(ch) {
		return true
	}

	if !Word(lastCh) && Word(ch) {
		return true
	}

	return false
}

// Increment each element in `vec` between `beg` and `end` by `INC`.
func IncVec(vec []int, inc *int, beg *int, end *int) []int {
	_inc := 1
	if inc != nil {
		_inc = *inc
	}
	_beg := 0
	if beg != nil {
		_beg = *beg
	}
	_end := len(vec)
	if end != nil {
		_end = *end
	}

	for _beg < _end {
		vec[_beg] += _inc
		_beg += 1
	}

	return vec
}

func GetHashForString(str string) map[int][]int {
	var result map[int][]int = make(map[int][]int)

	var strLen int = len(str)
	var index int = strLen - 1
	var runes = []rune(str)
	var downCh rune

	for 0 <= index {
		ch := runes[index]

		if Capital(&ch) {
			dictInsert(result, int(ch), index)

			downCh = toLower(ch)
		} else {
			downCh = ch
		}

		dictInsert(result, int(downCh), index)

		index -= 1
	}

	return result
}

type Result struct {
	indices []int
	score   int
	tail    int
}

// Return best score matching `query` against `str`.
func Score(str string, query string) *Result {
	if len(str) == 0 || len(query) == 0 {
		return nil
	}

	return &Result{}
}
