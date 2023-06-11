package credentials_test

import (
	"reflect"
	"testing"

	"github.com/srlk/go-commons/gcloud/credentials"
	"google.golang.org/api/option"
)

func TestNewCredentials(t *testing.T) {
	type args struct {
		credentials string
	}
	tests := []struct {
		name    string
		args    args
		want    option.ClientOption
		wantErr bool
	}{
		{
			name: "empty credentials fail",
			args: args{
				credentials: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non json string credentials fail",
			args: args{
				credentials: "{ i-am-not-json",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := credentials.NewCredentials(tt.args.credentials)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
