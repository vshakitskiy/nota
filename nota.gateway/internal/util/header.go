package util

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func IsHeaderAllowed(s string) (string, bool) {
	if _, ok := allowedHeaders[s]; ok {
		return s, true
	}

	return s, false
}
