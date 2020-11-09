package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandomNum(max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max)
}

func RandomNumStr(max int) string {
	fmtStr := fmt.Sprintf("0%dv", len(strconv.Itoa(max))-1)
	randomNum := RandomNum(max)
	return fmt.Sprintf("%"+fmtStr, randomNum)
}
