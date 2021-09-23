package utils

import (
	"strconv"
	"strings"
)

// DurToSec ...
func DurToSec(dur string) (sec float64) {
	durAry := strings.Split(dur, ":")
	if len(durAry) != 3 {
		return
	}
	hr, _ := strconv.ParseFloat(durAry[0], 64)
	sec = hr * (60 * 60)
	min, _ := strconv.ParseFloat(durAry[1], 64)
	sec += min * (60)
	second, _ := strconv.ParseFloat(durAry[2], 64)
	sec += second
	return
}
