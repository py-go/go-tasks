package main

import (
	"reflect"
	"testing"
)

// Estimate: 5mins
// Real: 5mins
func Test_testValidity(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "valid-alpha-numeric",
			str:  "23-ab-48-caba-56-haha",
			want: true,
		},
		{
			name: "invalid-input-format",
			str:  "23#ab-48-caba-56-haha",
			want: false,
		},

		{
			name: "invalid-input",
			str:  "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.str); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Estimate: 1mins
// Real: 1mins
func Test_averageNumber(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want float64
	}{
		{
			name: "valid-alpha-numeric",
			str:  "23-ab-48-caba-56-haha",
			want: float64(23+48+56) / float64(3),
		},
		{
			name: "valid-number",
			str:  "22",
			want: 22,
		},
		{
			name: "invalid-number",
			str:  "ab",
			want: 0,
		},
		{
			name: "invalid-input",
			str:  "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageNumber(tt.str); got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Estimate: <5mins
// Real: <5mins
func Test_wholeStory(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "valid1",
			str:  "1-hello-2-world",
			want: "hello world",
		},
		{
			name: "valid2",
			str:  "23-hello",
			want: "hello",
		},
		{
			name: "error",
			str:  "23",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wholeStory(tt.str); got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Estimate:5mins
// Real: 10mins
func Test_storyStats(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want *Stats
	}{
		{
			name: "success",
			str:  "23-hello-99-worldzz-333-aaabbb",
			want: &Stats{
				shortestWord:      "hello",
				longestWord:       "worldzz",
				avgWordLength:     6,
				avgWordLengthList: []string{"aaabbb"},
			},
		},
		{
			name: "success2",
			str:  "23-hello-99-worldzz-333",
			want: &Stats{
				shortestWord:  "hello",
				longestWord:   "worldzz",
				avgWordLength: 6,
			},
		},
		{
			name: "success3",
			str:  "23-hello-99-world-333",
			want: &Stats{
				shortestWord:      "hello",
				longestWord:       "hello",
				avgWordLength:     5,
				avgWordLengthList: []string{"hello", "world"},
			},
		},
		{
			name: "success4",
			str:  "23",
			want: &Stats{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storyStats(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storyStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
