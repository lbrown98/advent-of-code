package dec1

import "testing"

func Test_Distance_(t *testing.T) {
	tests := []struct {
		name   string
		team_a []int
		team_b []int
		want   int
	}{
		{
			name: "success",
			team_a: []int{3, 4, 2, 1, 3, 3},
			team_b: []int{4, 3, 5, 3, 9, 3},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := distance_sum(tt.team_a, tt.team_b)

			if ans != tt.want{
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}