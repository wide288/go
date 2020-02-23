package datetime

import "testing"

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