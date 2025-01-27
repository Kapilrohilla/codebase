package utils

func HandlePagination(page int, limit int) (int, int) {
	var finalPage int = page
	var finalLimit int = limit
	if page < 1 {
		finalPage = 1
	}

	if limit < 1 {
		finalLimit = 10
	}

	var offset int = (finalPage - 1) * finalLimit
	return offset, finalLimit
}
