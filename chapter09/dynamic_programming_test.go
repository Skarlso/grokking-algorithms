package chapter09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsack(t *testing.T) {
	items := []Item{
		{Name: "guitar", Value: 1500, Weight: 1},
		{Name: "stereo", Value: 3000, Weight: 4},
		{Name: "laptop", Value: 2000, Weight: 3},
	}
	assert.Equal(t, 3500, Knapsack(items, 4))
	items = append(items, Item{
		Name:   "iphone",
		Value:  2000,
		Weight: 1,
	})
	assert.Equal(t, 4000, Knapsack(items, 4))
}

func TestLongestCommonSubstring(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hish vs vista",
			args: args{
				a: "hish",
				b: "vista",
			},
			want: "is",
		},
		{
			name: "fish vs hish",
			args: args{
				a: "fish",
				b: "hish",
			},
			want: "ish",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LongestCommonSubstring(tt.args.a, tt.args.b), "LongestCommonSubstring(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fosh vs fort",
			args: args{
				a: "fosh",
				b: "fort",
			},
			want: 2,
		},
		{
			name: "fosh vs fish",
			args: args{
				a: "fosh",
				b: "fish",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LongestCommonSubsequence(tt.args.a, tt.args.b), "LongestCommonSubsequence(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
