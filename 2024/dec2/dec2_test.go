package dec2

import "testing"

func Test_Distance(t *testing.T) {
	tests := []struct {
		name   string
		want   int
	}{
		{
			name:   "success",
			want:   ??,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ??

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}