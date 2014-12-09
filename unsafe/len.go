package unsafe

import "example.com/repo.git/generic"

func Len(b generic.Value) int {
	return len(b)
}
