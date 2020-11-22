package hanoi

import "testing"

func setupWinState() {
	for i := 0 ; i < numDisks ; i++ {
		rods[0][i] = 0
		rods[1][i] = 0
		rods[2][i] = numDisks - i
	}
}

func TestCheckWinState_win(t *testing.T) {
	// setup
	setupWinState()

	// test
	got := checkWinState();
	if got != true {
		t.Errorf("checkWinState() = %t; want %t", got, true)
	}
}

func TestCheckWinState_not(t *testing.T) {
	// setup
	setupWinState()
	moveDisk("2", "1")

	// test
	got := checkWinState();
	if got != false {
		t.Errorf("checkWinState() = %t; want %t", got, false)
	}
}
