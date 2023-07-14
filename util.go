package main

import (
	"encoding/json"
	"strings"
)

func find[T any](xs []T, p func(v T) bool) bool {
	for _, x := range xs {
		if p(x) {
			return true
		}
	}
	return false
}

func contains[T comparable](xs []T, v T) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

func containsAll[T comparable](xs []T, ys ...T) bool {
	for _, y := range ys {
		if !contains(xs, y) {
			return false
		}
	}
	return true
}

func containsAny[T comparable](xs []T, ys ...T) bool {
	for _, y := range ys {
		if contains(xs, y) {
			return true
		}
	}
	return false
}

// hasPrefixFrom tests whether the string s begins with any prefix in ps.
func hasPrefixFrom(s string, ps []string) bool {
	for _, p := range ps {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

func JSON(v any) string {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error"
	}
	return string(j)
}
