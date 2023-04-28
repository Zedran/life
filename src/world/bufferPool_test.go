package world

import (
	"fmt"
	"testing"
)

/* Tests the pointer switch in BufferPool.NextState. Fails on improper pointer assignment. */
func TestBufferPoolNextState(t *testing.T) {
	bp := NewBufferPool(16)

	pState  := fmt.Sprintf("%p", bp.GetCurrentState() )
	pBuffer := fmt.Sprintf("%p", bp.GetCurrentBuffer())
	pSpare  := fmt.Sprintf("%p", bp.GetSpareBuffer()  )

	bp.NextState()

	if pSpare != fmt.Sprintf("%p", bp.GetCurrentBuffer()) {
		t.Error("Spare != Buffer")
	}

	if pBuffer != fmt.Sprintf("%p", bp.GetCurrentState()) {
		t.Error("Buffer != State")
	}

	if pState != fmt.Sprintf("%p", bp.GetSpareBuffer()) {
		t.Error("Current != Spare")
	}
}
