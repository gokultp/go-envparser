package envdecoder

import "testing"

type mockDecoder struct{}

func (md *mockDecoder) DecodeEnv() error { return nil }

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		val     interface{}
		wantErr bool
	}{
		{
			name:    "Wrong val type passed. Should throw error",
			val:     nil,
			wantErr: true,
		},
		{
			name:    "Correct val type passed. Should not throw error",
			val:     &mockDecoder{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decode(tt.val); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
