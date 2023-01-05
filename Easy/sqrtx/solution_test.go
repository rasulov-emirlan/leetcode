package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
