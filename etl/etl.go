// Package etl implements the "ETL" exercise.
package etl

import "strings"

// Transform transforms the given map into another one with a different format
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for score, letters := range in {
		for _, l := range letters {
			out[strings.ToLower(string(l))] = score
		}
	}
	return out
}
