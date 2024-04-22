package util

import (
	"fmt"
	"strconv"
)

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func Int64ToString(i int64) string {
	return fmt.Sprint(i)
}
