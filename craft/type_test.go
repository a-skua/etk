package craft

import (
	"fmt"
	"testing"
)

func TestMarginAll(t *testing.T) {
	tests := []struct {
		margin Margin
		want   Margin
	}{
		{
			MarginAll(10),
			Margin{10, 10, 10, 10},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			if tt.margin != tt.want {
				t.Errorf("MarginAll should return %v, but got %v", tt.want, tt.margin)
			}
		})
	}
}
