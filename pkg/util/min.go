package util

func MinInt64(args ...int64) int64 {
	min := args[0]
	for i := 1; i < len(args); i++ {
		if min > args[i] {
			min = args[i]
		}
	}
	return min
}
