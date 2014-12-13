// Package kind provides the Kind function for a generic.Value.
package kind

import "github.com/anlhord/generic"

// Kind returns 's' when the generic.Value is a Slice, 'i' when it's interface.
func Kind(b generic.Value) byte {
	if cap(b) == 0 && len(b) > 0 {
		return 'i'
	}
	return 's'
}
