package request

import (
	"testing"
)

func TestPost(t *testing.T) {
	type args struct {
		url         string
		contentType string
		body        []byte
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
				url:         BASE_URL + "/post",
				contentType: "application/json",
				body:        []byte(`{"message": "hello world"}`),
			},
			want:    200,
			wantErr: false,
		},
		{
			name: "500",
			args: args{
				url:         BASE_URL + "/status/500",
				contentType: "application/json",
				body:        []byte(`{"message": "hello world"}`),
			},
			want:    500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := Post(tt.args.url, tt.args.contentType, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("Post() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
