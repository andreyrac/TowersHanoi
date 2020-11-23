package hanoi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// setupWinState set up the disk on the rods in a winning state (end of game state)
func setupWinState() {
	for i := 0; i < numDisks; i++ {
		rods[0][i] = 0
		rods[1][i] = 0
		rods[2][i] = numDisks - i
	}
}

// setupBlank set up the rods to have no disks at all
func setupBlank() {
	for i := 0; i < numDisks; i++ {
		rods[0][i] = 0
		rods[1][i] = 0
		rods[2][i] = 0
	}
}

// TestCheckWinState_win checks that winning condition is detected properly
func TestCheckWinState_win(t *testing.T) {
	// setup
	setupWinState()

	// test
	actual := checkWinState()
	expected := true
	require.Equal(t, actual, expected, "checkWinState() actual[%t] expected[%t]", actual, expected)
}

// TestCheckWinState_not checks that non-winning condition is detected properly
func TestCheckWinState_not(t *testing.T) {
	// setup
	setupWinState()
	moveDisk("2", "1")

	// test
	actual := checkWinState()
	expected := false
	require.Equal(t, actual, expected, "checkWinState() actual[%t] expected[%t]", actual, expected)
}

// TestValidMove_toEmptyRod tests moving a disk to an empty rod
func TestValidMove_toEmptyRod(t *testing.T) {
	// setup
	Restart()

	// test
	actual := moveDisk("0", "1")
	expected := ValidMove
	require.Equal(t, actual, expected, "moveDisk() actual[%d] expected[%d]", actual, expected)
	actual = rods[0][numDisks-1]
	expected = 0
	require.Equal(t, actual, expected, "moveDisk() didn't remove disk from original position actual[%d] expected[%d]", actual, expected)

	actual = rods[1][0]
	expected = 1
	require.Equal(t, actual, expected, "moveDisk() didn't place disk in correct 'to' position actual[%d] expected[%d]", actual, expected)
}

// TestValidMove_toNonEmptyRod tests moving a disk to a non-empty rod
func TestValidMove_toNonEmptyRod(t *testing.T) {
	// setup
	setupBlank()
	rods[0][0] = 1
	rods[0][1] = 2
	rods[1][0] = 3

	// test
	actual := moveDisk("0", "1")
	expected := ValidMove

	require.Equal(t, actual, expected, "moveDisk() actual[%d] expected[%d]", actual, expected)

	actual = rods[0][1]
	expected = 0
	require.Equal(t, actual, expected, "moveDisk() didn't remove disk from original position actual[%d] expected[%d]", actual, expected)

	actual = rods[1][1]
	expected = 2
	require.Equal(t, actual, expected, "moveDisk() didn't place disk in correct 'to' position actual[%d] expected[%d]", actual, expected)
}

// TestValidMove_fromAlmostEmptyRod tests moving the last disk from a rod
func TestValidMove_fromAlmostEmptyRod(t *testing.T) {
	// setup
	setupBlank()
	rods[0][0] = 1
	rods[1][0] = 2

	// test
	actual := moveDisk("0", "1")
	expected := ValidMove
	require.Equal(t, actual, expected, "moveDisk() actual[%d] expected[%d]", actual, expected)

	actual = rods[0][0]
	expected = 0
	require.Equal(t, actual, expected, "moveDisk() didn't remove disk from original position actual[%d] expected[%d]", actual, expected)

	actual = rods[1][1]
	expected = 1
	require.Equal(t, actual, expected, "moveDisk() didn't place disk in correct 'to' position actual[%d] expected[%d]", actual, expected)
}

// TestInvalidMove_fromEmptyRod tests an invalid move where the 'from' rod is empty
func TestInvalidMove_fromEmptyRod(t *testing.T) {
	// setup
	Restart()

	// test
	actual := moveDisk("1", "2")
	expected := InvalidMove
	require.Equal(t, actual, expected, "moveDisk() actual[%d] expected[%d]", actual, expected)
}

// TestInvalidMove_ontoSmallerDisk ensures you cannot move a larger disk onto a smaller disk
func TestInvalidMove_ontoSmallerDisk(t *testing.T) {
	// setup
	setupBlank()
	rods[0][0] = 2
	rods[1][0] = 1

	// test
	actual := moveDisk("0", "1")
	expected := InvalidMove
	require.Equal(t, actual, expected, "moveDisk() actual[%d] expected[%d]", actual, expected)

	// ensure rewind worked
	actual = rods[0][0]
	expected = 2
	require.Equal(t, actual, expected, "moveDisk() rewind failed actual[%d] expected[%d]", actual, expected)
}

// TestInvalidMove_fromInvalidRod tests range checking on 'from'
func TestInvalidMove_fromInvalidRod(t *testing.T) {
	// setup
	Restart()

	// test lower bound
	actual := moveDisk("-1", "1")
	expected := InvalidMove
	require.Equal(t, actual, expected, "moveDisk(-1, 1) actual[%d] expected[%d]", actual, expected)

	// test upper bound
	actual = moveDisk("3", "1")
	require.Equal(t, actual, expected, "moveDisk(3, 1) actual[%d] expected[%d]", actual, expected)
}

// TestInvalidMove_toInvalidRod tests range checking on 'to'
func TestInvalidMove_toInvalidRod(t *testing.T) {
	// setup
	Restart()

	// test lower bound
	actual := moveDisk("-1", "2")
	expected := InvalidMove
	require.Equal(t, actual, expected, "moveDisk(1, -1) actual[%d] expected[%d]", actual, expected)

	// test upper bound
	actual = moveDisk("0", "3")
	require.Equal(t, actual, expected, "moveDisk(0, 3) actual[%d] expected[%d]", actual, expected)
}

// TestInvalidMove_invalidFromRodString tests handling on bad input for 'from'
func TestInvalidMove_invalidFromRodString(t *testing.T) {
	// setup
	Restart()

	// test
	actual := moveDisk("a", "2")
	expected := InvalidMove
	require.Equal(t, actual, expected, "moveDisk(a, 2) actual[%d] expected[%d]", actual, expected)
}

// TestInvalidMove_invalidToRodString tests handling bad input for 'to'
func TestInvalidMove_invalidToRodString(t *testing.T) {
	// setup
	Restart()

	// test
	actual := moveDisk("0", "f")
	expected := InvalidMove
	require.Equal(t, actual, expected, "moveDisk(0, f) actual[%d] expected[%d]", actual, expected)
}

// TestWinningMove ensures we detect win state upon making final move
func TestWinningMove(t *testing.T) {
	// setup
	setupWinState()
	rods[1][0] = 1
	rods[2][numDisks-1] = 0

	// test
	actual := moveDisk("1", "2")
	expected := WinningMove
	require.Equal(t, actual, expected, "moveDisk() actual[%d] expected[%d]", actual, expected)
}
