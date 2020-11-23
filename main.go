package main

import (
	"github.com/andreyrac/TowersHanoi/hanoi"
	"net/http"
)

func main() {
	hanoi.Restart()
	http.HandleFunc("/state", hanoi.PostState)
	http.HandleFunc("/move", hanoi.MoveDisk)
	http.HandleFunc("/hasWon", hanoi.HasWon)
	http.ListenAndServe(":5051", nil)
}
