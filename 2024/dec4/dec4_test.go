package dec4

import "testing"

func Test_findHorizontal(t *testing.T) {
	tests := []struct {
		name    string
		input 	string
		want    int
	}{
		{
			name: "forward simple",
			input: "MMMSXXMASM",
			want: 1,
		},
		{
			name: "backwards simple",
			input: "MMMSXSAMXM",
			want: 1,
		},
		{
			name: "combo",
			input: "XMASAMXAMM",
			want: 2,
		},
		{
			name: "forward edge",
			input: "XMASXXASXMAS",
			want: 2,
		},
		{
			name: "backward edge",
			input: "SAMXAAASSAMX",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indices := searchForLetter(tt.input, 'X')
			ans := checkHorizontal(tt.input, indices)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}

func Test_findVertical(t *testing.T) {
	tests := []struct {
		name    string
		input 	[]string
		row_number int
		want    int
	}{
		{
			name: "backward simple",
			input: []string{
				"MMMSXXMASM", "MSAMXMSMSA", "AMXSXMAAMM", 
				"MSAMASMSMX", "XMASAMXAMM", "XXAMMXXAMA",
			},
			row_number: 4,
			want: 1,
		}, 
		{
			name: "forward simple",
			input: []string{
				"XMASAMXAMM", "MXAMMXXAMA", "AMSMSASXSS", 
				"SAXAMASAAA","MAMMMXMMMM",
			},
			row_number: 0,
			want: 1,
		},
		{
			name: "forward edge",
			input: []string{
				"XMASAMXAMM", "MXAMMXXAMA", "AMSMSASXSS", 
				"SAXAMASAAA",
			},
			row_number: 0,
			want: 1,
		},	
		{
			name: "backward edge",
			input: []string{
				"MMMSXXMASM", "MSAMXMSMSA", "AMXSXMAAMM", 
				"MSAMASMSMX", "XMASAMXAMM",
			},
			row_number: 4,
			want: 1,
		}, 	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indices := searchForLetter(tt.input[tt.row_number], 'X')
			ans := checkVertical(tt.input, tt.row_number, indices)

			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}