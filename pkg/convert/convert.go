package convert

import (
	"encoding/json"
	"fmt"
	"math"
)

func MoneyToMinor(v float32) int64 {
	return int64(math.Round(float64(v * 100.0)))
}

func MinorToMoney(v int64) float32 {
	return float32(v) / 100
}

func MinorToNumber(v int64) json.Number {
	vStr := ""
	switch {
	case v < 10:
		vStr = fmt.Sprintf("0.0%d", v)
	case v < 100:
		vStr = fmt.Sprintf("0.%d", v)
	default:
		vStr = fmt.Sprintf("%d.%d", v/100, v%100)
	}
	return json.Number(vStr)
}
