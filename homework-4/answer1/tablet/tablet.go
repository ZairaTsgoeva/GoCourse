package tablet

import "fmt"

type tablet struct {
	locked bool
	charge float32
}

func CreateTablet(charge float32) *tablet {
	return &tablet{charge: charge, locked: true}
}

func (t *tablet) checkAccess() bool {
	return true //todo
}

func (t *tablet) checkCharge() bool {
	return t.charge > 0
}

func (t *tablet) lock() {
	t.locked = true
}

func (t *tablet) isLocked() bool {
	return t.locked
}

func (t *tablet) unlock() bool {
	success := t.checkAccess()
	if success {
		t.locked = false
	}
	return success
}

func (t *tablet) Open() bool {
	return t.checkCharge() && t.unlock()
}

func (t *tablet) Read() bool {
	success := !t.isLocked()
	if success {
		t.charge -= 5
	}
	return success
}

func (t *tablet) Close() {
	t.lock()
}

func (t *tablet) State() string {
	return fmt.Sprintf("tablet - locked: %v, charge: %f", t.locked, t.charge)
}