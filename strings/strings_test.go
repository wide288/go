package strings

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				s: "hello world",
				substr: "llw",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				s: "hello world",
				substr: "ll",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				str: " 1space ",
			},
			want: "1space",
		},
		{
			name: "normal",
			args: args{
				str: "  2space3   ",
			},
			want: "2space3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSpace(tt.args.str); got != tt.want {
				t.Errorf("TrimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimRight(t *testing.T) {
	type args struct {
		str    string
		cutset string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				str: "hello,",
				cutset: ",",
			},
			want: "hello",
		},
		{
			name: "normal",
			args: args{
				str: "hello,,,",
				cutset: ",",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimRight(tt.args.str, tt.args.cutset); got != tt.want {
				t.Errorf("TrimRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitInt(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				str: "1,2,3,4",
				sep: ",",
			},
			want: []int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitInt(tt.args.str, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		strArr []string
		sep    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "join strings",
			args: args{
				strArr: []string{"ab", "bc", "cd", "de"},
				sep:    ",",
			},
			want: "ab,bc,cd,de",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.strArr, tt.args.sep); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}