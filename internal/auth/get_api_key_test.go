package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}
	tests := []test{
		{name: "no header", headers: nil, want: "", wantErr: true},
		{name: "no auth header", headers: http.Header{"test": []string{"123"}}, want: "", wantErr: true},
		{name: "auth header wrong", headers: http.Header{"Authorization": []string{"123"}}, want: "", wantErr: true},
		{name: "correct", headers: http.Header{"Authorization": []string{"ApiKey XYZ"}}, want: "XYZ", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
