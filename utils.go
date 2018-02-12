package ling

type is func(rune) bool

func StringIs(s string, f is) bool {
	for _, c := range s {
		if !f(c) {
			return false
		}
	}
	return true
}
