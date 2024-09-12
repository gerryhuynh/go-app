package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Player struct {
	Name string
	Wins int
}

type League []Player

func main() {
	path := os.Args[1]

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(fmt.Errorf("problem opening %s, %v", path, err))
	}
	defer file.Close()

	league, err := NewLeague(file)
	if err != nil {
		log.Fatal(fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err))
	}

	fmt.Println(league)
}

func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
