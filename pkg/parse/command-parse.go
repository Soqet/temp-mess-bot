package parse

func DeleteFirstWord(source string) string {
	for i := uint64(0); i < uint64(len(source)); i++ { 
		if source[i] == ' ' {
			return source[i + 1:]
		}
	}
	return ""
}

func DeleteFirstNWords(source string, n uint64) string {
	spaceCounter := uint64(0)
	for i := uint64(0); i < uint64(len(source)); i++ { 
		if source[i] == ' ' {
			spaceCounter++
			if spaceCounter >= n {
				return source[i + 1:]
			}
		}
	}
	return ""
}

func ParseWord(source string, n uint64) string {
	var count uint64 = 0
	var start uint64 = 0 // [ .. )
	var end uint64 = 0
	for i := uint64(0); i < uint64(len(source)); i++ {
		if source[i] == ' ' {
			count++
			if count == n {
				start = i + 1
			}
			if count == n + 1 {
				end = i
			}
		}
	}
	if start != 0 {
		if end != 0 {
			return source[start:end]
		}
		return source[start:]
	}
	return ""
}