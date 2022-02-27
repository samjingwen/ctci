package qns1_3

import (
	"testing"
)

func TestURLify(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Mr John Smith"}, "Mr%20John%20Smith"},
		{"2", args{"12 ⌘3"}, "12%20⌘3"},
		{"2", args{"abc123"}, "abc123"},
		{"2", args{" abc123"}, "%20abc123"},
		{"2", args{" abc123 "}, "%20abc123%20"},
	}
	for _, tt := range tests {
		got := &tt.args.url
		t.Run(tt.name, func(t *testing.T) {
			URLify(&tt.args.url)
			if *got != tt.want {
				t.Errorf("URLify() = %v, want %v", *got, tt.want)
			}
		})
	}
}
