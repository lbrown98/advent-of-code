package dec2

import "testing"

func Test_SafeReports(t *testing.T) {
	tests := []struct {
		name    string
		reports [][]int
		want    int
	}{
		{
			name: "success",
			reports: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := SafeReports(tt.reports)

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
				// {7, 6, 4, 2, 1},
				// {1, 2, 7, 8, 9},
				// {9, 7, 6, 2, 1},
				// {1, 3, 2, 4, 5},
				// {8, 6, 4, 4, 1},
				// {1, 3, 6, 7, 9},
				{4, 3, 5, 7, 9},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Dampener(tt.reports)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}
