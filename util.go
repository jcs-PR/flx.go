/**
 * $File: util.go $
 * $Date: 2024-12-22 09:37:57 $
 * $Revision: $
 * $Creator: Jen-Chieh Shen $
 * $Notice: See LICENSE.txt for modification and distribution information
 *                   Copyright © 2024 by Shen, Jen-Chieh $
 */

package flx

// Return true if `slice` contains the `item`.
func contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
