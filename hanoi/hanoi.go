package hanoi

import (
	"encoding/json"
	"fmt"
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
			return false
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
	// cheap shortcut, assumes moveDisk is working correctly
	return rods[2][numDisks-1] != 0
}

// moveDisk internal call used by MoveDisk.
// returns InvalidMove, ValueMove, or WinningMove
func moveDisk(fromS string, toS string) int {
	// input 'from' value
	from, errF := strconv.Atoi(fromS)
	if errF != nil {
		fmt.Printf("failed to convert 'from' string[%s]: %s\n", fromS, errF)
		return InvalidMove
	}
	if from < 0 || from >= 3 {
		fmt.Printf("'from' value out of range: %d\n", from)
		return InvalidMove
	}

	// input 'to' value
	to, errT := strconv.Atoi(toS)
	if errT != nil {
		fmt.Printf("failed to convert 'to' string[%s]: %s\n", toS, errT)
		return InvalidMove
	}
	if to < 0 || to >= 3 {
		fmt.Printf("'to' value out of range: %d\n", to)
		return InvalidMove
	}

	// make the move, return early on invalid moves
	disk := getTopDisk(&rods[from])
	if disk == 0 {
		fmt.Printf("'from' rod %d is empty\n", from)
		return InvalidMove
	}
	if !moveDiskTo(&rods[to], disk) {
		fmt.Printf("cannot stack larger disk[%d] onto smaller disk on rod %d\n", disk, to)
		moveDiskTo(&rods[from], disk) // rewind move
		return InvalidMove
	}

	// send the valid move or winning move response
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

// MoveDisk performs movement of disks and responds with: InvalidMove, ValidMove or WinningMove
func MoveDisk(w http.ResponseWriter, r *http.Request) {
	response := moveDisk(r.FormValue("From"), r.FormValue("To"))
	json.NewEncoder(w).Encode(response)
}

// HasWon responds with true if state is a winning state
func HasWon(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(checkWinState())
}
