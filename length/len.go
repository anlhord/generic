// Package length provides the Len function for a generic.Value
package length

import "github.com/anlhord/generic"

func Len(b generic.Value) int {
	if cap(b) == 0 && len(b) > 0 {
		return 0
	}
	return len(b)
}
