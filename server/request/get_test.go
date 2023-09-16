package request

import (
	"testing"
)

var BASE_URL = "https://httpbin.org"

func TestGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "200",
			args: args{
				url: BASE_URL + "/get",
			},
			want:    200,
			wantErr: false,
		},
		{
			name: "500",
			args: args{
				url: BASE_URL + "/status/500",
			},
			want:    500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
