package service

import "testing"

func TestCheckUserModeAndFetchUserInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckUserModeAndFetchUserInfo(); (err != nil) != tt.wantErr {
				t.Errorf("CheckUserModeAndFetchUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
