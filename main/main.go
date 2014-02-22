package main

import (
	"encoding/json"
	"fmt"
	. "github.com/humbhenri/tic_tac_toe"
	"io/ioutil"
	"net/http"
)

type End struct {
	Winner Mark
    Board Board
}

func Json(b Board) (s string) {
	js, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		s = ""
		return
	}
	s = string(js)
	return
}

func swapPlayer(player Mark) Mark {
	if player == X {
		return O
	} else if player == O {
		return X
	}
	return None
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

func sendJson(w http.ResponseWriter, st interface{}) {
	json, err := json.Marshal(st)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	fmt.Println(string(json))
}

func gameEnded(b *Board) (bool, Mark) {
	end := false
	win := b.Win()
	if win != None || b.FreePositions() == 0 {
		end = true
	}
	return end, win
}

func sendEndMessageIfGameOver(w http.ResponseWriter, b *Board) bool {
	end, mark := gameEnded(b)
	if end {
		sendJson(w, End{mark, *b})
		return true
	}
	return false
}

func main() {
	b := &Board{}
	b.Start()
	player := X

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
		err = Play(b, player)
		if err != nil {
			panic(err)
		} else {
			sendJson(w, *b)
		}
		if sendEndMessageIfGameOver(w, b) {
			return
		}
		player = swapPlayer(player)
		fmt.Println(b.String())

	})

	http.ListenAndServe(":8080", nil)
}
