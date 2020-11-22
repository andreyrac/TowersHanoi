package hanoi

import (
	"testing"
	"strconv"
)

func setupBlank() {
	for i := 0; i < numDisks; i++ {
		rods[0][i] = 0
		rods[1][i] = 0
		rods[2][i] = 0
	}
}


func TestValidMove_toEmptyRod(t *testing.T) {
	// setup
	Restart();

	// test
	got := moveDisk("0", "1")
	if got != ValidMove {
		t.Errorf("moveDisk(\"0\",\"1\") = %d; want %d", got, ValidMove)
	}
	if rods[0][numDisks - 1] != 0 {
		t.Errorf("moveDisk didn't remove disk from original position: %d", rods[0][numDisks - 1])
	}
	if rods[1][0] != 1 {
		t.Errorf("moveDisk didn't place disk in correct 'to' position: %d", rods[1][0]);
	}
}

func TestValidMove_toNonEmptyRod(t *testing.T) {
	// setup
	setupBlank();
	rods[0][0] = 1
	rods[0][1] = 2
	rods[1][0] = 3

	// test
	got := moveDisk("0", "1");
	if got != ValidMove {
		t.Errorf("moveDisk(\"0\",\"1\") = %d; want %d", got, ValidMove)
	}
	if rods[0][1] != 0 {
		t.Errorf("moveDisk didn't remove disk from original position: %d", rods[0][1])
	}
	if rods[1][1] != 2 {
		t.Errorf("moveDisk didn't place disk in correct 'to' position: %d", rods[1][1]);
	}
}

func TestValidMove_fromAlmostEmptyRod(t *testing.T) {
	// setup
	setupBlank();
	rods[0][0] = 1
	rods[1][0] = 2

	// test
	got := moveDisk("0", "1");
	if got != ValidMove {
		t.Errorf("moveDisk(\"0\",\"1\") = %d; want %d", got, ValidMove)
	}
	if rods[0][0] != 0 {
		t.Errorf("moveDisk didn't remove disk from original position: %d", rods[0][0])
	}
	if rods[1][1] != 1 {
		t.Errorf("moveDisk didn't place disk in correct 'to' position: %d", rods[1][1]);
	}

}

func TestInvalidMove_fromEmptyRod(t *testing.T) {
	// setup
	Restart();

	// test
	got := moveDisk("1", "2")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"1\",\"2\") = %d; want %d", got, InvalidMove)
	}
}

func TestInvalidMove_ontoSmallerDisk(t *testing.T) {
	// setup
	setupBlank();
	rods[0][0] = 2
	rods[1][0] = 1

	// test
	got := moveDisk("0", "1")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"0\",\"1\") = %d; want %d", got, InvalidMove)
	}

}

func TestInvalidMove_fromInvalidRod(t *testing.T) {
	// setup
	Restart();

	// test lower bound
	got := moveDisk("-1", "1")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"-1\",\"1\") = %d; want %d", got, InvalidMove)
	}

	// test upper bound
	from := strconv.Itoa(numDisks)
	got = moveDisk(from, "1")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"%s\",\"1\") = %d; want %d", from, got, InvalidMove)
	}
}

func TestInvalidMove_toInvalidRod(t *testing.T) {
	// setup
	Restart();

	// test lower bound
	got := moveDisk("-1", "2")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"-1\",\"2\") = %d; want %d", got, InvalidMove)
	}

	to := strconv.Itoa(numDisks)
	got = moveDisk("0", to)
	if got != InvalidMove {
		t.Errorf("moveDisk(\"0\",\"%s\") = %d; want %d", to, got, InvalidMove)
	}
}

func TestInvalidMove_invalidFromRodString(t *testing.T) {
	// setup
	Restart();

	// test
	got := moveDisk("a", "2")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"a\",\"2\") = %d; want %d", got, InvalidMove)
	}
}

func TestInvalidMove_invalidToRodString(t *testing.T) {
	// setup
	Restart();

	// test
	got := moveDisk("0", "f")
	if got != InvalidMove {
		t.Errorf("moveDisk(\"0\",\"f\") = %d; want %d", got, InvalidMove)
	}
}

func TestWinningMove(t *testing.T) {
	// setup
	for i := 0 ; i < numDisks ; i++ {
		rods[0][i] = 0
		rods[1][i] = 0
		rods[2][i] = numDisks - i
	}
	rods[1][0] = 1
	rods[2][numDisks - 1] = 0

	// test
	got := moveDisk("1", "2")
	if got != WinningMove {
		t.Errorf("moveDisk(\"1\",\"2\") = %d; want %d", got, WinningMove)
	}

}
