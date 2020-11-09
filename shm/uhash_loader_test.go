package shm

import "testing"

func TestLoadHash(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadHash(); (err != nil) != tt.wantErr {
				t.Errorf("LoadHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
