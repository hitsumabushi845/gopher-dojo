package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6) + 1

	switch s {
	case 6:
		println(s, "大吉")
	case 5, 4:
		println(s, "中吉")
	case 3, 2:
		println(s, "吉")
	case 1:
		println(s, "凶")
	}
}
