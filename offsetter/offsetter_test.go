package offsetter_test

import (
	"testing"

	"github.com/wasanx25/goss/event"
	"github.com/wasanx25/goss/offsetter"
)

func TestOffsetterUpdateAndGet(t *testing.T) {
	tests := []struct {
		eventType event.Type
		offset    int
		max       int
		limit     int
		expected  int
	}{
		{event.PageDown, 1, 10, 1, 1 + 1},
		{event.PageDown, 8, 10, 1, 8 + 1},
		{event.PageDown, 12, 10, 1, 12},
		{event.PageDown, 1, 10, 11, 1},
		{event.PageUp, 10, 1, 1, 9},
		{event.PageDownHalf, 1, 10, 6, 1 + 3},
		{event.PageDownHalf, 7, 10, 2, 7 + 1},
		{event.PageDownHalf, 12, 10, 1, 12},
		{event.PageDownHalf, 1, 10, 11, 1},
		{event.PageUpHalf, 1, 10, 6, 1},
		{event.PageUpHalf, 7, 10, 2, 7 - 1},
		{event.PageUpHalf, 3, 10, 8, 1},
		{event.PageDownScreen, 1, 10, 3, 4},
		{event.PageDownScreen, 3, 10, 6, 4},
		{event.PageDownScreen, 5, 10, 5, 5},
		{event.PageUpScreen, 1, 10, 6, 1},
		{event.PageUpScreen, 7, 10, 2, 7 - 2},
		{event.PageUpScreen, 3, 10, 8, 1},
		{event.PageEnd, 10, 3, 2, 3 - 1},
		{event.PageTop, 10, 0, 0, 0},
	}

	for _, tt := range tests {
		o := offsetter.NewOffsetter(tt.offset, tt.max, tt.limit)
		actual := o.UpdateAndGet(tt.eventType)

		if actual != tt.expected {
			t.Errorf("eventType=%v, expected=%d, got=%d", tt.eventType, tt.expected, actual)
		}
	}
}
