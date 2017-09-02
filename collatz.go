package main

import (
	"html/template"
	"log"
	"math/big"
	"net/http"
)

var big_one = big.NewInt(1)
var big_three = big.NewInt(3)

func collatz_step(a *big.Int) {
	if a.Bit(0) == 0 {
		a.Rsh(a, 1)
	} else {
		a.Mul(a, big_three)
		a.Add(a, big_one)
	}
}

func collatz_step_count_maximum(next_step, max, counter *big.Int) {
	n := big.NewInt(0)
	n.Set(next_step)
	collatz_step(next_step)
	max.Set(n)
	for n.Cmp(big_one) > 0 {
		collatz_step(n)
		if n.Cmp(max) > 0 {
			max.Set(n)
		}
		counter.Add(counter, big_one)
	}
}

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.New("").ParseGlob("./templates/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", page_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type PageData struct {
	Current   *big.Int
	Steps     *big.Int
	Maximum   *big.Int
	NextStep  *big.Int
	ReachDown *big.Int
	ReachUp   *big.Int
	AtOne     bool
}

func page_handler(w http.ResponseWriter, r *http.Request) {
	var number *big.Int
	if r.URL.Path == "/" {
		number = big.NewInt(1)
	} else {
		number = big.NewInt(0)
		_, success := number.SetString(r.URL.Path[1:], 10)
		if !success {
			w.WriteHeader(404)
			w.Write([]byte("404 page not found"))
			return
		}
	}

	data := PageData{
		Current:   number,
		Steps:     big.NewInt(0),
		Maximum:   big.NewInt(0),
		NextStep:  big.NewInt(0).Set(number),
		ReachDown: nil,
		ReachUp:   nil,
		AtOne:     number.Cmp(big_one) == 0,
	}
	collatz_step_count_maximum(data.NextStep, data.Maximum, data.Steps)

	temp_div := big.NewInt(0)
	temp_mod := big.NewInt(0)
	temp_div.Sub(number, big_one)
	temp_div.DivMod(temp_div, big_three, temp_mod)
	if temp_mod.Cmp(big.NewInt(0)) == 0 && temp_div.Cmp(big_one) > 0 {
		data.ReachDown = temp_div
	}
	data.ReachUp = big.NewInt(0).Lsh(number, 1)

	err := tmpl.ExecuteTemplate(w, "index.tmpl", data)
	if err != nil {
		log.Println(err)
	}
}
