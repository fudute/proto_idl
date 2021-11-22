package main

import (
	"testing"
)

func TestExtracProtocGoPackage(t *testing.T) {
	type args struct {
		proto string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{args: args{proto: "../service/echoservice.proto"}, want: "github.com/fudute/echoservice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtracProtocGoPackage(tt.args.proto)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtracProtocGoPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtracProtocGoPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}
