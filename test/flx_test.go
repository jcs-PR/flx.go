/**
 * $File: flx_test.go $
 * $Date: 2024-12-22 09:38:50 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package flx_test

import "testing"
import "github.com/the-flx/flx.go"

func TestWord1(t *testing.T) {
	ch := 'c'
	result := flx.Word(&ch)
	if result == true {
		t.Errorf("Test `Word` 1: %v", result)
	}
}

func TestWord2(t *testing.T) {
	ch := ' '
	result := flx.Word(&ch)
	if result == false {
		t.Errorf("Test `Word` 2: %v", result)
	}
}
