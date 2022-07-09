package main

import (
	"math/rand"
	"time"
)

var lowerCaseRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var mixedCaseRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var upperCaserRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int, letter string) string {
	var letterRunes []rune
	b := make([]rune, n)

	switch letter {
	case "lower":
		letterRunes = lowerCaseRunes
	case "mixed":
		letterRunes = mixedCaseRunes
	case "upper":
		letterRunes = upperCaserRunes
	default:
		letterRunes = lowerCaseRunes
	}

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NextAlias(last string) string {
	if last == "" {
		return "a"
	} else if last[len(last)-1] == 'z' {
		return last[:len(last)-1] + "aa"
	} else {
		return last[:len(last)-1] + string(last[len(last)-1]+1)
	}
}
