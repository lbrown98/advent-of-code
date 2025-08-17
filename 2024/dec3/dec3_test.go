package dec3

import (
	"testing"
)

func Test_matchMultiplies(t *testing.T) {
	tests := []struct {
		name    string
		input 	string
		want    int
	}{
		{
			name: "success",
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := matchMultiplies(tt.input)

			if len(ans) != tt.want {
				t.Errorf("want %d, got %d", tt.want, len(ans))
			}
		})
	}
}

func Test_calculation(t *testing.T) {
	tests := []struct {
		name    string
		input 	[]string
		want    int
	}{
		{
			name: "success",
			input: []string{
				"mul(2,4)",
				"mul(5,5)",
				"mul(11,8)",
				"mul(8,5)",
			},
			want: 161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := calculation(tt.input)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}

func Test_matchMultipliesAndInstructions(t *testing.T) {
	tests := []struct {
		name    string
		input 	string
		want    int
	}{
		{
			name: "success",
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := matchMultipliesAndInstructions(tt.input)

			if len(ans) != tt.want {
				t.Errorf("want %d, got %d", tt.want, len(ans))
			}
		})
	}
}

func Test_calculationWithInstructions(t *testing.T) {
	tests := []struct {
		name    string
		input 	[]string
		want    int
	}{
		{
			name: "success",
			input: []string{
				"mul(2,4)",
				"don't()",
				"mul(5,5)",
				"mul(11,8)",
				"do()",
				"mul(8,5)",
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := calculationWithInstructions(tt.input)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}