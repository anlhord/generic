package unsafe

import "github.com/anlhord/generic"

func Len(b generic.Value) int {
	if cap(b) == 0 && len(b) > 0 {
		return 1
	}
	return len(b)
}
