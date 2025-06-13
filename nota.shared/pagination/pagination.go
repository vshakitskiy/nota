package pagination

import "math"

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func TotalPages(total int64, size int) int {
	return int(math.Ceil(float64(total) / float64(size)))
}
