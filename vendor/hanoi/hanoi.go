package hanoi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const numDisks = 4

// Responses to /move
const (
	InvalidMove int = iota
	ValidMove
	WinningMove
)

// rods Response to /state and holding variable of the actual state
var rods = [3][numDisks]int{}

// getTopDisk removes the top disk from rod.
// returns the disc number or 0 if no disk present
func getTopDisk(rod *[numDisks]int) int {
	for i := numDisks - 1; i >= 0; i-- {
		if rod[i] != 0 {
			var ret = rod[i]
			rod[i] = 0
			return ret
		}
	}
	// no disk on this rod
	return 0
}

// moveDiskTo moves disk to top of the top of this rod.
// returns true if move is valid
func moveDiskTo(rod *[numDisks]int, disk int) bool {
	for i := numDisks - 2; i >= 0; i-- {
		if rod[i] == 0 {
			continue
		}
		if rod[i] < disk {
			// attempting to stack larger disk on smaller disk
			return false;
		}
		// top of stack found and valid
		rod[i+1] = disk
		return true
	}
	// rod has no disk
	rod[0] = disk
	return true
}

// checkWinState checks the rods to see if the last rod holds all the disks
func checkWinState() bool {
	// check for winning condition
	for i := 0; i < numDisks; i++ {
		if rods[2][i] != numDisks - i {
			return false
		}
	}
	return true
}

// moveDisk internal call used by MoveDisk
func moveDisk(fromS string, toS string) int {
	from, errF := strconv.Atoi(fromS)
	to, errT := strconv.Atoi(toS)
	if errF != nil || errT != nil || from < 0 || from >= numDisks || to < 0 || to >= numDisks {
		return InvalidMove
	}

	// make the move, return early on invalid moves
	disk := getTopDisk(&rods[from])
	if disk == 0 {
		return InvalidMove
	}
	if !moveDiskTo(&rods[to], disk) {
		moveDiskTo(&rods[from], disk) // rewind move
		return InvalidMove
	}

	// send the valid move or winning move response
	var response = ValidMove;
	if checkWinState() {
		return WinningMove
	} else {
		return ValidMove
	}
}

// Restart restarts the state of the rods to the beginning
func Restart() {
	for i := 0; i < numDisks; i++ {
		rods[0][i] = numDisks - i
		rods[1][i] = 0
		rods[2][i] = 0
	}
}

// PostState posts the state of the Hanoi Tower Rods
func PostState(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rods)
}

// MoveDisk performs movement of disks and responds with: InvalidMod, ValidMove or WinningMove
func MoveDisk(w http.ResponseWriter, r *http.Request) {
	response := moveDisk(r.formValue("From"), r.FormValue("To");
	json.NewEncoder(w).Encode(response)
}

// HasWon responds with true if state is a winning state
func HasWon(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(checkWinState())
}

// RestartState restarts the game to the beginning state and posts the state as the response
func RestartState(w http.ResponseWriter, r *http.Request) {
	Restart();
	PostState(w, r);
}

