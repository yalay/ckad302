package controllers

import (
	"strconv"
	"strings"
)

// Atoi convert string to int
func Atoi(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

func Atoi32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}
