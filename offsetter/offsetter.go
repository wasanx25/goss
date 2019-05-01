package offsetter

import (
	"github.com/wasanx25/goss/event"
)

type Offsetter interface {
	UpdateAndGet(e event.Type) int
}

type offsetter struct {
	offset int
	max    int
	limit  int
}

func NewOffsetter(offset, max, limit int) Offsetter {
	return &offsetter{
		offset: offset,
		max:    max,
		limit:  limit,
	}
}

func (o *offsetter) UpdateAndGet(e event.Type) int {
	switch e {
	case event.PageDown:
		o.pageDown()
	case event.PageUp:
		o.pageUp()
	case event.PageDownHalf:
		o.pageDownHalf()
	case event.PageUpHalf:
		o.pageUpHalf()
	case event.PageDownScreen:
		o.pageDownWindow()
	case event.PageUpScreen:
		o.pageUpWindow()
	case event.PageEnd:
		o.pageEnd()
	case event.PageTop:
		o.pageTop()
	}

	return o.offset
}

func (o *offsetter) pageDown() {
	if o.max < o.limit {
		return
	}

	if o.limit+o.offset < o.max {
		o.offset++
	}
}

func (o *offsetter) pageUp() {
	if o.offset > 0 {
		o.offset--
	}
}

func (o *offsetter) pageDownHalf() {
	if o.max > o.offset+o.limit/2*3 {
		o.offset = o.offset + o.limit/2
	} else if o.offset+o.limit < o.max {
		o.offset = o.max - o.limit
	}
}

func (o *offsetter) pageUpHalf() {
	if o.offset > o.limit/2 {
		o.offset = o.offset - o.limit/2
	} else if o.offset > 1 && o.offset < o.limit/2 {
		o.offset = 1
	}
}

func (o *offsetter) pageDownWindow() {
	if o.max > o.offset+o.limit*2 {
		o.offset = o.offset + o.limit
	} else if o.offset+o.limit < o.max {
		o.offset = o.max - o.limit
	}
}

func (o *offsetter) pageUpWindow() {
	if o.offset > o.limit {
		o.offset = o.offset - o.limit
	} else if o.offset > 1 && o.offset < o.limit {
		o.offset = 1
	}
}

func (o *offsetter) pageEnd() {
	if o.max > o.limit {
		o.offset = o.max - o.limit + 1
	}
}

func (o *offsetter) pageTop() {
	o.offset = 0
}
