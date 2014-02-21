package main

import (
	"fmt"
	. "github.com/humbhenri/tic_tac_toe"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type End struct {
    winner Mark
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
		var cell interface{}
		err = json.Unmarshal(body, &cell)
		if err != nil {
			panic(err)
		}

		pos := cell.(map[string]interface{})["pos"]
		row := pos.([]interface{})[0]
		col := pos.([]interface{})[1]
		b.Put(player, int(row.(float64)), int(col.(float64)))
        player = swapPlayer(player)
        err = Play(b, player)
        win := b.Win()
        if err != nil {
            fmt.Println(err)
        } else {
            if win == None {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprint(w, Json(*b))
            } else {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprint(w, End{win})
            }

        }
        player = swapPlayer(player)
            
        fmt.Println(b.String())
	})

	http.ListenAndServe(":8080", nil)
}
