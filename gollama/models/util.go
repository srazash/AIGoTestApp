package models

func ToKB(b int) float64 {
	return float64(b) / 1000
}

func ToMB(b int) float64 {
	return float64(b) / 1000 / 1000
}

func ToGB(b int) float64 {
	return float64(b) / 1000 / 1000 / 1000
}
