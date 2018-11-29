package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ESBN   string `json:"esbn"`
}

func (b Book) ToJSON() []byte {
	ToJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJson
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	Book{Title: "Go Lang", Author: "P M F", ESBN: "ESBN-2353-1422"},
	Book{Title: "Go Lang2", Author: "W P M F", ESBN: "ESBN-2353-1429"},
}

func BooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)

}
