package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(100000) + 899999)
}
