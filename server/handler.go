package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/mxngls/kled-server/parser"
)

func Search(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()

	word := q.Get("word")
	lang := q.Get("lang")
	langCode := q.Get("langCode")
	page := q.Get("page")

	resp := requestSearch(word, lang, langCode, page)

	defer resp.Body.Close()

	data, err := parser.ParseSearch(resp.Body, langCode)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	JSON, err := JSONMarshal(data)
	if err != nil {
		panic(err)
	}

	w.Write(JSON)
}

func View(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()

	id := q.Get("id")
	lang := q.Get("lang")
	langCode := q.Get("langCode")

	resp := requestView(id, lang, langCode)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(body)

	data, err := parser.ParseView(reader, id, langCode)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	JSON, err := JSONMarshal(data)
	if err != nil {
		panic(err)
	}

	w.Write(JSON)
}
