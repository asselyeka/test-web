package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func mainPage(w http.ResponseWriter, r *http.Request) {

	body := r.FormValue("body")
	font := r.FormValue("fonts")
	if font == "" {
		font = "standard"
	}
	data := asciiArt(body, font)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func main() {

	http.HandleFunc("/", mainPage)

	port := ":8080"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Server", err)
	}

}

func asciiArt(str string, filename string) []string {

	//str := os.Args[1]

	// read from file
	fileName := filename + ".txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")
	var finalString []string
	//print string to terminal
	for h := 0; h < 9; h++ {
		for _, l := range str {

			for i, line := range lines {
				if i == (int(l)-32)*9+h {
					finalString = append(finalString, line)
				}
			}

		}
		finalString = append(finalString, "\n")
	}

	return finalString
}
