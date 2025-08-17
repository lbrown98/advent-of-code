package dec2

import (
	"slices"
	"testing"
	"github.com/lbrown98/advent-of-code/2024/config"
)

func Test_SafeReports(t *testing.T) {
	tests := []struct {
		name    string
		reports [][]int
		want    int
	}{
		{
			name: "success",
			reports: [][]int{
				{7, 6, 4, 2, 1}, //safe
				{1, 2, 7, 8, 9}, //unsafe
				{9, 7, 6, 2, 1}, //unsafe
				{1, 3, 2, 4, 5}, //unsafe
				{8, 6, 4, 4, 1}, //unsafe
				{1, 3, 6, 7, 9}, //safe
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SafeReports(tt.reports, false)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}

func Test_Dampener(t *testing.T) {
	tests := []struct {
		name    string
		reports [][]int
		want    int
	}{
		{
			name: "success",
			reports: [][]int{
				{7, 6, 4, 2, 1}, //safe
				{1, 2, 7, 8, 9}, //unsafe
				{9, 7, 6, 2, 1}, //unsafe
				{1, 3, 2, 4, 5}, //safe
				{8, 6, 4, 4, 1}, //safe
				{1, 3, 6, 7, 9}, //safe
				{4, 3, 5, 7, 9}, //safe
				{1, 3, 4, 18, 19}, //unsafe
				{1, 4, 4, 5, 6}, //safe
				{1, 3, 7, 6, 6}, //unsafe
				{1, 3, 7, 5, 6}, //safe
				{71, 69, 70, 71, 72, 75}, //safe
			},
			want: 8,
		},
		{
			name: "real examples",
			reports: config.Levels,
			want:381,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SafeReports(tt.reports, true)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}

func Test_ModifyReport(t *testing.T) {
	tests := []struct {
		name    string
		reports []int
		removeIndex    int
		want []int
	}{
		{
			name: "remove index 0",
			reports: []int{7, 6, 4, 3, 1},
			removeIndex: 0,
			want: []int{6, 4, 3, 1},
		},
		{
			name: "remove index 1",
			reports: []int{7, 6, 4, 3, 1},
			removeIndex: 1,
			want: []int{7, 4, 3, 1},
		},
		{
			name: "remove index 4",
			reports: []int{7, 6, 4, 3, 1},
			removeIndex: 3,
			want: []int{7, 6, 4, 1},
		},
		{
			name: "remove final index",
			reports: []int{7, 6, 4, 3, 1},
			removeIndex: 4,
			want: []int{7, 6, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := removeIndex(tt.reports, tt.removeIndex)

			if !slices.Equal(ans, tt.want) {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}