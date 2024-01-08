package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	m := InitMonster()
	res := m.Store()
	if !res {
		t.Fatalf("Test Store() failed: %v", res)
	}
	t.Logf("Test Store() success: %v", res)
}

func TestRestore(t *testing.T) {
	m := Monster{}
	res := m.Restore()
	if !res {
		t.Fatalf("Test Store() failed: %v", res)
	}
	t.Logf("Test Store() success: %v", res)
}
