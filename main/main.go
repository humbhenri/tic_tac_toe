package main

import (
	"fmt"
	. "github.com/humbhenri/tic_tac_toe"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
	b := &Board{}
	b.Start()
	//player := X

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

		mark := cell.(map[string]interface{})["mark"]
		pos := cell.(map[string]interface{})["pos"]
		row := pos.([]interface{})[0]
		col := pos.([]interface{})[1]
		fmt.Println(mark, row.(float64), col.(float64))

	})

	http.ListenAndServe(":8080", nil)
}
