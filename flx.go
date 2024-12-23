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
	"math"
	"reflect"
	"slices"
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
		return true
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

// Return map for string where keys are characters.
// Value is a sorted list of indexes for character occurrences.
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

// Generate the heatmap vector of string.
//
// See documentation for logic.
func GetHeatmapStr(str string, groupSeparator *rune) []int {
	strLen := len(str)
	strLastIndex := strLen - 1

	var scores []int = fillSlice(strLen, defaultScore)

	var penaltyLead rune = '.'

	// final char bonus
	scores[strLastIndex] += 1

	var groupAlist [][]int = [][]int{{-1, 0}}

	// Establish baseline mapping
	var lastCh *rune = nil
	var groupWordCount int
	var index1 int = 0

	for _, ch := range str {
		// before we find any words, all separaters are
		// considered words of length 1.  This is so "foo/__ab"
		// gets penalized compared to "foo/ab".
		var effectiveLastChar *rune = nil
		if groupWordCount != 0 {
			effectiveLastChar = lastCh
		}

		if Boundary(effectiveLastChar, &ch) {
			groupAlist[0] = slices.Insert(groupAlist[0], 2, index1)
		}

		if !Word(lastCh) && Word(&ch) {
			groupWordCount += 1
		}

		// ++++ -45 penalize extension
		if lastCh != nil && *lastCh == penaltyLead {
			scores[index1] += -45
		}

		if groupSeparator != nil && groupSeparator == &ch {
			groupAlist[0][1] = groupWordCount
			groupWordCount = 0
			groupAlist = slices.Insert(groupAlist, 0, []int{index1, groupWordCount})
		}

		if index1 == strLastIndex {
			groupAlist[0][1] = groupWordCount
		} else {
			lastCh = &ch
		}

		index1 += 1
	}

	var groupCount int = len(groupAlist)
	var separatorCount int = groupCount - 1

	// ++++ slash group-count penalty
	if separatorCount != 0 {
		var val int = groupCount * -2
		scores = IncVec(scores, &val, nil, nil)
	}

	var index2 int = separatorCount
	var lastGroupLimit *int = nil
	var basepathFound bool = false

	// score each group further
	for _, group := range groupAlist {
		var groupStart int = group[0]
		var wordCount int = group[1]
		// this is the number of effective word groups
		var wordsLength int = len(group) - 2
		var basepathP bool = false

		if wordsLength != 0 && !basepathFound {
			basepathFound = true
			basepathP = true
		}

		var num int

		if basepathP {
			// ++++ basepath separator-count boosts
			var boost int = 0
			if separatorCount > 1 {
				boost = separatorCount - 1
			}
			// ++++ basepath word count penalty
			var penalty = -wordCount
			num = 35 + boost + penalty
		} else { // ++++ non-basepath penalties
			if index2 == 0 {
				num = -3
			} else {
				num = -5 + (index2 - 1)
			}
		}

		beg := (groupStart + 1)
		scores = IncVec(scores, &num, &beg, lastGroupLimit)

		var cddrGroup []int = sliceCopy(group) // clone

		cddrGroup = slices.Delete(cddrGroup, 0, 2)

		var wordIndex int = wordsLength - 1
		var lastWord int = strLen
		if lastGroupLimit != nil {
			lastWord = *lastGroupLimit
		}

		for _, word := range cddrGroup {
			// ++++  beg word bonus AND
			scores[word] += 85

			var index3 int = word
			var charI int = 0

			for index3 < lastWord {
				scores[index3] += (-3 * wordIndex) - charI

				charI += 1
				index3 += 1
			}

			lastWord = word
			wordIndex -= 1
		}

		lastGroupLimit = new(int)
		*lastGroupLimit = groupStart + 1

		index2 -= 1
	}

	return scores
}

// Return sublist bigger than `val` from sorted `sorted`.
func BiggerSublist(sorted []int, val *int) []int {
	if sorted == nil || val == nil {
		return sorted
	}

	var result []int = []int{}

	for _, sub := range sorted {
		if sub > *val {
			result = append(result, sub)
		}
	}

	return result
}

// Recursively compute the best match for a string, passed as `strInfo` and
// `heatmap`, according to `query`.
func FindBestMatch(imatch *[]Result, strInfo map[int][]int, heatmap []int, greaterThan *int, query string, queryLen int, qIndex int, matchCache *map[int][]Result) {
	var greaterNum int = 0
	if greaterThan != nil {
		greaterNum = *greaterThan
	}
	var hashKey int = qIndex + (greaterNum * queryLen)
	var hashVal []Result = dictGet(matchCache, &hashKey)

	if hashVal != nil {
		clear(*imatch)
		for _, val := range hashVal {
			*imatch = append(*imatch, val)
		}
	} else {
		var uchar int = int(strAt(query, qIndex))
		var sorted = dictGet(&strInfo, &uchar)
		var indexes []int = BiggerSublist(sorted, greaterThan)
		var tempScore int
		var bestScore int = math.MinInt32

		if qIndex >= queryLen-1 {
			// At the tail end of the recursion, simply generate all possible
			// matches with their scores and return the list to parent.
			for _, val := range indexes {
				var indices []int = []int{}
				indices = append(indices, val)
				*imatch = append(*imatch, Result{indices, heatmap[val], 0})
			}
		} else {
			for _, val := range indexes {
				var elemGroup []Result = []Result{}

				var clonedStrInfo map[int][]int = dictCopy(strInfo) // clone
				var clonedHeatmap []int = sliceCopy(heatmap)        // clone

				FindBestMatch(&elemGroup, clonedStrInfo, clonedHeatmap, &val, query, queryLen, qIndex+1, matchCache)

				for _, elem := range elemGroup {
					if reflect.DeepEqual(elem, Result{}) {
						continue
					}

					var caar int = elem.Indices[0]
					var cadr int = elem.Score
					var cddr int = elem.Tail

					if (caar - 1) == val {
						tempScore = cadr + heatmap[val] + (min(cddr, 3) * 15) + 60
					} else {
						tempScore = cadr + heatmap[val]
					}

					// We only care about the optimal match, so only forward the match
					// with the best score to parent
					if tempScore > bestScore {
						bestScore = tempScore

						clear(*imatch)
						var indices []int = sliceCopy(elem.Indices)
						indices = append([]int{val}, indices...)
						var tail int = 0
						if (caar - 1) == val {
							tail = cddr + 1
						}
						*imatch = append(*imatch, Result{indices, tempScore, tail})
					}
				}
			}
		}

		// Calls are cached to avoid exponential time complexity
		dictSet(matchCache, &hashKey, sliceCopy(*imatch))
	}
}

type Result struct {
	Indices []int
	Score   int
	Tail    int
}

// Return best score matching `query` against `str`.
func Score(str string, query string) *Result {
	if len(str) == 0 || len(query) == 0 {
		return nil
	}

	var strInfo map[int][]int = GetHashForString(str)
	var heatmap = GetHeatmapStr(str, nil)

	var queryLen = len(query)
	var fullMatchBoost bool = (1 < queryLen) && (queryLen < 5)
	var matchCache map[int][]Result = map[int][]Result{}
	var optimalMatch []Result = []Result{}
	FindBestMatch(&optimalMatch, strInfo, heatmap, nil, query, queryLen, 0, &matchCache)

	if len(optimalMatch) == 0 {
		return nil
	}

	var result *Result = &optimalMatch[0]
	var caar int = len(result.Indices)

	if fullMatchBoost && caar == len(str) {
		result.Score += 10000
	}

	return result
}
