package maths_test

import (
	"reflect"
	"testing"

	"github.com/srlk/go-commons/maths"
)

func Test_Max(t *testing.T) {
	type testCase struct {
		name string
		a    int
		b    int
		want int
	}
	cases := []testCase{
		{
			name: "int a < b",
			a:    1,
			b:    2,
			want: 2,
		},
		{
			name: "float a < b",
			a:    1.0,
			b:    2.0,
			want: 2.0,
		},
		{
			name: "float a > b",
			a:    11.0,
			b:    2.0,
			want: 11.0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maths.Max(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Min(t *testing.T) {
	type testCase struct {
		name string
		a    int
		b    int
		want int
	}
	cases := []testCase{
		{
			name: "int a < b",
			a:    1,
			b:    2,
			want: 1,
		},
		{
			name: "float a < b",
			a:    1.0,
			b:    2.0,
			want: 1.0,
		},
		{
			name: "float a > b",
			a:    11.0,
			b:    2.0,
			want: 2.0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maths.Min(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
