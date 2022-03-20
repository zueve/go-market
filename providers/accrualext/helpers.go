package accrualext

import "math"

func MoneyToMinor(v float32) int64 {
	return int64(math.Round(float64(v * 100.0)))
}

func MinorToMoney(v int64) float32 {
	return float32(v) / 100
}
