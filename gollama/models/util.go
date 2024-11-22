package models

func ToKB(b int) float64 {
	return float64(b) / 1024
}

func ToMB(b int) float64 {
	return float64(b) / 1024 / 1024
}

func ToGB(b int) float64 {
	return float64(b) / 1024 / 1024 / 1024
}
