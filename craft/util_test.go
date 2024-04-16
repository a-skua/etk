package craft

import (
	"fmt"
	"testing"
)

func Test_calcSize(t *testing.T) {
	tests := []struct {
		size   Size
		margin Margin
		want   Size
	}{
		{Size{100, 50}, Margin{}, Size{100, 50}},
		{Size{100, 50}, Margin{10, 0, 10, 0}, Size{120, 50}},
		{Size{100, 50}, Margin{0, 10, 0, 10}, Size{100, 70}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			size := calcSize(tt.size, tt.margin)
			if size != tt.want {
				t.Errorf("calcSize should return %v, but got %v", tt.want, size)
			}
		})
	}
}
