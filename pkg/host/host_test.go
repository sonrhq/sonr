package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	t.Run("Should Create object, state should be IDLE", func(t *testing.T) {
		fsm := NewFSM(nil)
		assert.NotNil(t, fsm)
		assert.NotNil(t, fsm.Current)
		assert.Equal(t, fsm.Current, Status_IDLE)
	})
	t.Run("Should report correct state on transition", func(t *testing.T) {
		fsm := NewFSM(nil)
		fsm.SetStatus(Status_STANDBY)
		assert.Equal(t, fsm.Current, Status_STANDBY)
	})
	t.Run("Should not transition on invalid state transition", func(t *testing.T) {
		fsm := NewFSM(nil)
		fsm.SetStatus(Status_READY)
		assert.Equal(t, fsm.Current, Status_IDLE)
	})
	t.Run("Operation Pause Should destroy channel", func(t *testing.T) {
		fsm := NewFSM(nil)
		fsm.PauseOperation()
		assert.NotNil(t, fsm.Chn)
		assert.Equal(t, fsm.flag, uint64(1))
	})

}
