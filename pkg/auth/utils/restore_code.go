package utils

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

func GenerateRestoreCode() string {
	code := strings.Builder{}
	for i := 0; i < 6; i++ {
		num := strconv.FormatUint(rand.Uint64N(10), 10)
		code.WriteString(num)
	}

	return code.String()
}

func CompareRestoreCode(codeInput string, codeCache string) bool {
	if codeInput != codeCache {
		return false
	}

	return true
}
