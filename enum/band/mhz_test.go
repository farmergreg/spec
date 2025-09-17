package band

import "testing"

func TestMHz_Format(t *testing.T) {
	tests := []struct {
		name string
		mhz  MHz
		want string
	}{
		{"simple", 14.070, "14.07"},
		{"no decimal", 7, "7"},
		{"many decimals", 3.567890, "3.56789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mhz.String(); got != tt.want {
				t.Errorf("MHz.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
