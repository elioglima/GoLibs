package libs

func In(value int, i ...int) bool {
	for _, index := range i {
		if value == index {
			return true
		}
	}

	return false
}
