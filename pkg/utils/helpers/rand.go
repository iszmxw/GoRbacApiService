package helpers

import (
	"math/rand"
	"strconv"
	"time"
)

func Rand(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := ""
	for index := 0; index < num; index++ {
		randNum += strconv.Itoa(r.Intn(10))
	}
	return randNum
}
