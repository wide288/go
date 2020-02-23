package datetime

import (
	"fmt"
	"testing"
)

func TestTimeUnixToStr(t *testing.T) {
	type args struct {
		unix int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "unix2string",
			args: args{
				unix: 1582466531,
			},
			want: "2020-02-23 10:02:11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeUnixToStr(tt.args.unix); got != tt.want {
				t.Errorf("TimeUnixToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

/**
cur month time : 2020-2-23 22:09:30
return:
	got: 2020-02-01 12:00:00
	got1: 2020-03-01 12:00:00
 */
func TestGetLastMonth(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
	}{
		// TODO: Add test cases.
		{
			name:  "1 month",
			args:  args{
				num: 1,
			},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetLastMonth(tt.args.num)
			fmt.Println("got:", TimeUnixToStr(got))
			fmt.Println("got1:", TimeUnixToStr(got1))

			if got != tt.want {
				t.Errorf("GetLastMonth() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLastMonth() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}