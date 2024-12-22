/**
 * $File: flx.go $
 * $Date: 2024-12-22 07:52:08 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package flx

var wordSeparators = []rune{' ', '-', '_', ':', '.', '/', '\\'}

// Check if `ch` is a word character.
func Word(ch *rune) bool {
	if ch == nil {
		return false
	}
	return contains(wordSeparators, *ch)
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
