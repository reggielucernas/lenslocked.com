package main

import (
	"html/template"
	"os"
	"time"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	mp := make(map[string]string)
	//mp["hometown"] = "Vancouver"
	mp["current"] = "Seattle"

	data := struct {
		Name   string
		Date   string
		Cities map[string]string
	}{"John Smith", time.Now().Format("2006-01-02"), mp}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
