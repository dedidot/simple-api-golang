package utils

import (
	"github.com/dedidot/generate/stringer"
	"strconv"
)

func GenerateId() int {
	generateResult := stringer.RandomStr(5, "nozero")
	generateInt, err := strconv.Atoi(generateResult)
	if err != nil {
		return 0
	}
	return generateInt
}
