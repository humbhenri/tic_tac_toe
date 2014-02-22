package main

import (
	"encoding/json"
	"fmt"
	"github.com/humbhenri/tic_tac_toe"
	"io/ioutil"
	"net/http"
)

type message struct {
	Winner tic_tac_toe.Mark
	Board  tic_tac_toe.Board
}

func swapPlayer(player tic_tac_toe.Mark) tic_tac_toe.Mark {
	if player == tic_tac_toe.X {
		return tic_tac_toe.O
	} else if player == tic_tac_toe.O {
		return tic_tac_toe.X
	}
	return tic_tac_toe.None
}

func getMarkFromPlayer(body []byte) (int, int) {
	var cell interface{}
	err := json.Unmarshal(body, &cell)
	if err != nil {
		panic(err)
	}

	pos := cell.(map[string]interface{})["pos"]
	row := pos.([]interface{})[0]
	col := pos.([]interface{})[1]
	return int(row.(float64)), int(col.(float64))
}

func sendJSON(w http.ResponseWriter, st interface{}) {
	json, err := json.Marshal(st)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	fmt.Println(string(json))
}

func gameEnded(b *tic_tac_toe.Board) (bool, tic_tac_toe.Mark) {
	end := false
	win := b.Win()
	if win != tic_tac_toe.None || b.FreePositions() == 0 {
		end = true
	}
	return end, win
}

func sendEndMessageIfGameOver(w http.ResponseWriter, b *tic_tac_toe.Board) bool {
	end, mark := gameEnded(b)
	if end {
		sendJSON(w, message{mark, *b})
		return true
	}
	return false
}

func main() {
	b := &tic_tac_toe.Board{}
	b.Start()
	player := tic_tac_toe.X

	fmt.Println("Listening on port 8080.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/mark", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			panic(err)
		}

		row, col := getMarkFromPlayer(body)
		fmt.Println(row, col)
		b.Put(player, row, col)
		if sendEndMessageIfGameOver(w, b) {
			return
		}
		player = swapPlayer(player)
		err = tic_tac_toe.Play(b, player)
		if err != nil {
			panic(err)
		}
		if sendEndMessageIfGameOver(w, b) {
			return
		}
		sendJSON(w, message{tic_tac_toe.None, *b})
		player = swapPlayer(player)
		fmt.Println(b.String())

	})

	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		b.Start()
		player = tic_tac_toe.X
	})

	http.ListenAndServe(":8080", nil)
}
