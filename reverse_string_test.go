package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOutput string
	}{
		{
			name:       "should reverse \"Hello, world!\" string",
			input:      "Hello, world!",
			wantOutput: "!dlrow ,olleH",
		}, {
			name: "should reverse multiline string",
			input: `Hello,
world!`,
			wantOutput: "!dlrow\n,olleH",
		}, {
			name:       "should reverse string with international characters",
			input:      `Hello, світ!`,
			wantOutput: "!тівс ,olleH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := ReverseString(tt.input); gotOutput != tt.wantOutput {
				t.Errorf("ReverseString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
