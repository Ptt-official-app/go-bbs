package shm

import "testing"

func TestLoadUHash(t *testing.T) {
	isTest = true

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadUHash(); (err != nil) != tt.wantErr {
				t.Errorf("loadUHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
