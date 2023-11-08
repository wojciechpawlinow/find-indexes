package memory

import (
	"context"
	"errors"

	"github.com/wojciechpawlinow/find-indexes/internal/domain/index"
)

type SliceRepository struct {
	Values *[]int
}

var _ index.Repository = (*SliceRepository)(nil)

// OneByValue returns an index of a given value or an index of the closest value within the boundaries of conformation
// v - value to find
// returns:
//
//	idx - index of the value,
//	directMatch - true if the value was found, false if the closest value was returned,
//	err - error message.
func (r *SliceRepository) OneByValue(_ context.Context, v int) (int, int, bool, error) {

	// so having a sorted list the efficient way to find an index for a value is to use binary search

	totalItems := len(*r.Values)
	if totalItems == 0 {
		return -1, -1, false, errors.New("empty slice")
	}

	lowIdx := 0
	highIdx := totalItems - 1
	closestIdx := -1

	// calculate the tolerance for the value
	conformation := v / 10
	minAcceptableValue := v - conformation
	maxAcceptableValue := v + conformation

	for lowIdx <= highIdx {

		// current index and value
		midIdx := (lowIdx + highIdx) / 2
		midValue := (*r.Values)[midIdx]

		if midValue == v {
			return midIdx, midValue, true, nil // exact match
		} else if midValue < minAcceptableValue {
			lowIdx = midIdx + 1
		} else if midValue > maxAcceptableValue {
			highIdx = midIdx - 1
		} else {

			// current value at midIdx is within the tolerance range, lets replace closestIdx with it
			closestIdx = midIdx

			// narrow the search
			if midValue < v {
				lowIdx = midIdx + 1
			} else {
				highIdx = midIdx - 1
			}
		}
	}

	// exact match not found, return the closest index
	if closestIdx != -1 {
		return closestIdx, (*r.Values)[closestIdx], false, nil
	}

	return -1, -1, false, nil
}
