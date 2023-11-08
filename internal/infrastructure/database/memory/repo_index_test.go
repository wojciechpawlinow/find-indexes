package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceRepository_OneByValue(t *testing.T) {

	tests := []struct {
		name           string
		values         []int
		searchValue    int
		wantIdx        int
		wantValue      int
		wantExactMatch bool
		wantErr        error
	}{
		{
			name:           "empty slice",
			values:         []int{},
			searchValue:    100,
			wantIdx:        -1,
			wantValue:      -1,
			wantExactMatch: false,
			wantErr:        errors.New("empty slice"),
		},
		{
			name:           "exact match",
			values:         []int{50, 100, 150, 200},
			searchValue:    150,
			wantIdx:        2,
			wantValue:      150,
			wantExactMatch: true,
			wantErr:        nil,
		},
		{
			name:           "within tolerance, lower bound",
			values:         []int{50, 100, 150, 200},
			searchValue:    145,
			wantIdx:        2,
			wantValue:      150,
			wantExactMatch: false,
			wantErr:        nil,
		},
		{
			name:           "within tolerance, upper bound",
			values:         []int{50, 100, 150, 200},
			searchValue:    155,
			wantIdx:        2,
			wantValue:      150,
			wantExactMatch: false,
			wantErr:        nil,
		},
		// Add more test cases as necessary...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SliceRepository{
				Values: &tt.values,
			}
			gotIdx, gotValue, gotExactMatch, err := r.OneByValue(context.Background(), tt.searchValue)

			if tt.wantErr != nil {
				assert.Error(t, err, "OneByValue() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.NoError(t, err, "OneByValue() unexpected error = %v", err)
				assert.Equal(t, tt.wantIdx, gotIdx, "OneByValue() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
				assert.Equal(t, tt.wantValue, gotValue, "OneByValue() gotValue = %v, want %v", gotValue, tt.wantValue)
				assert.Equal(t, tt.wantExactMatch, gotExactMatch, "OneByValue() gotExactMatch = %v, want %v", gotExactMatch, tt.wantExactMatch)
			}
		})
	}
}
