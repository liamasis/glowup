package ui

import (
	"slices"
	"strings"
)

func sortMarkdowns(mds []*markdown, sortOption string) {
	switch sortOption {
	case "alpha":
		slices.SortStableFunc(mds, func(a, b *markdown) int {
			return strings.Compare(a.Note, b.Note)
		})
	case "time":
		slices.SortStableFunc(mds, func(a, b *markdown) int {
			if a.Modtime.After(b.Modtime) {
				return -1
			}
			if b.Modtime.After(a.Modtime) {
				return 1
			}
			return 0
		})
	default:
		// Default to time
		slices.SortStableFunc(mds, func(a, b *markdown) int {
			if a.Modtime.After(b.Modtime) {
				return -1
			}
			if b.Modtime.After(a.Modtime) {
				return 1
			}
			return 0
		})
	}
}
