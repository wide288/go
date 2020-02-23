package converter

import (
	"reflect"
	"testing"
)

func TestStrToByteArray(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				str: "abcdef",
			},
			want: []byte{97, 98, 99, 100, 101, 102},
		},
		{
			name: "normal",
			args: args{
				str: "1234567890",
			},
			want: []byte{49, 50, 51, 52, 53, 54, 55, 56, 57, 48},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToByteArray(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteArrayToStr(t *testing.T) {
	type args struct {
		bA []byte
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
				[]byte{49, 50, 51, 52, 53, 54, 55, 56, 57, 48},
			},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteArrayToStr(tt.args.bA); got != tt.want {
				t.Errorf("ByteArrayToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}