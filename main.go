package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const numDisks = 4

// Move used for /move requests
type Move struct {
	From int `json:"From"`
	To   int `json:"To"`
}

// Responses to /move
const (
	InvalidMove int = iota
	ValidMove
	WinningMove
)

var rods = [3][numDisks]int{}

func restart() {
	for i := 0; i < numDisks; i++ {
		rods[0][i] = numDisks - i
		rods[1][i] = 0
		rods[2][i] = 0
	}
}

// PostState posts the state of the Hanoi Tower Rods
func postState(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rods)
}

func getTopDisk(rod [numDisks]int) int {
	for i := numDisks - 1; i >= 0; i-- {
		if rod[i] != 0 {
			ret := rod[i]
			rod[i] = 0
			return ret
		}
	}
	// no disk on this rod
	return 0
}

func moveDiskTo(rod [numDisks]int, disk int) int {
	for i := numDisks - 2; i >= 0; i-- {
		if rod[i] > disk {
			// top of stack found and valid
			rod[i+1] = disk
			return disk
		}
		if rod[i] != 0 {
			// attempting to stack larger disk on smaller disk
			return 0
		}
	}
	// rod has no disk
	rod[0] = disk
	return disk
}

// MoveDisk performs movement of disks and responds with: InvalidMod, ValidMove or WinningMove
func moveDisk(w http.ResponseWriter, r *http.Request) {
	//TODO: check that these are set
	from, err := strconv.ParseInt(r.FormValue("From"), 0, 8)
	to, err := strconv.ParseInt(r.FormValue("To"), 0, 8)

	//read movement value
	var m = Move{From: from, To: to}
	//json.NewDecoder(r.URL).Decode(&m)

	// make the move, return early on invalid moves
	disk := getTopDisk(rods[m.From])
	if disk == 0 {
		json.NewEncoder(w).Encode(InvalidMove)
		return
	}
	if moveDiskTo(rods[m.To], disk) == 0 {
		moveDiskTo(rods[m.From], disk) // rewind move
		json.NewEncoder(w).Encode(InvalidMove)
		return
	}

	// check for winning condition
	var response = WinningMove
	for i := 0; i < numDisks; i++ {
		if rods[2][i] != numDisks-1 {
			response = ValidMove
			break
		}
	}

	// send the valid move or winning move response
	json.NewEncoder(w).Encode(response)
}

func main() {
	restart()
	http.HandleFunc("/state", postState)
	http.HandleFunc("/move", moveDisk)
	http.ListenAndServe(":5051", nil)
}
