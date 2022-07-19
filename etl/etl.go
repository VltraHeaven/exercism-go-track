package etl

import (
    "strings"
)

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
    for i := range in {
        for _, v := range in[i] {
            out[strings.ToLower(v)] += i
        }
    }
	return out
}
