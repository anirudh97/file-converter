package utils

import (
	"reflect"
	"testing"
)

func TestDetectType(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Whitespace string",
			input: "   ",
			want:  "",
		},
		{
			name:  "Integer",
			input: "42",
			want:  42,
		},
		{
			name:  "Negative integer",
			input: "-123",
			want:  -123,
		},
		{
			name:  "Float",
			input: "3.14",
			want:  3.14,
		},
		{
			name:  "Negative float",
			input: "-2.718",
			want:  -2.718,
		},
		{
			name:  "Scientific notation",
			input: "1.23e5",
			want:  123000.0,
		},
		{
			name:  "True boolean",
			input: "true",
			want:  true,
		},
		{
			name:  "False boolean",
			input: "false",
			want:  false,
		},
		{
			name:  "True boolean with spaces",
			input: "  true  ",
			want:  true,
		},
		{
			name:  "True boolean uppercase",
			input: "TRUE",
			want:  true,
		},
		{
			name:  "False boolean uppercase",
			input: "FALSE",
			want:  false,
		},
		{
			name:  "Normal string",
			input: "hello",
			want:  "hello",
		},
		{
			name:  "String with spaces",
			input: "  hello world  ",
			want:  "hello world",
		},
		{
			name:  "String that looks numeric",
			input: "42abc",
			want:  "42abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectType(tt.input)

			// Check type and value
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("DetectType() type = %T, want type %T", got, tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectType() = %v, want %v", got, tt.want)
			}
		})
	}
}
