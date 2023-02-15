package main

func Cats(start, finish int) int {
	div := finish - start
	if div == 0 {
		return 0
	} else if div%2 == 0 {
		return 2
	} else if div%2 == 1 {
		return 1
	}
	return div
}
