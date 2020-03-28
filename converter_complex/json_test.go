package converter_complex

import (
	"reflect"
	"testing"
)

func Test_relation2json(t *testing.T) {
	type args struct {
		rel Diy
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test diy struct to json",
			args:    args{
				rel: Diy{
					"Hello Katu",
					"World",
					[]string{"111", "112", "113"},
					},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := relation2json(tt.args.rel)
			t.Log("got:", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("relation2json() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("relation2json() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_json2relation(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Diy
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TO struct",
			args:    args{
				data: []byte(`{"Test1":"Hello Katu","Test2":"World","Arrs":["111","112","113"]}`),
			},
			want:    Diy{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json2relation(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("json2relation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json2relation() got = %v, want %v", got, tt.want)
			}
		})
	}
}