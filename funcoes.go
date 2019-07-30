package GoLibs

import "strings"

func IfthenI(condicao bool, r1, r2 int) int {
	if condicao {
		return r1
	}

	return r2
}

func IfthenS(condicao bool, r1, r2 string) string {
	if condicao {
		return r1
	}

	return r2
}

func IfthenB(condicao bool, r1, r2 bool) bool {
	if condicao {
		return r1
	}

	return r2
}

func Len(v string) int {
	return len(strings.TrimSpace(v))
}
