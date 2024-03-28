package utils

func CompareRestoreCode(codeInput string, codeCache string) bool {
	if codeInput != codeCache {
		return false
	}

	return true
}
