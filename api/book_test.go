package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {
	book := Book{Title: "Go Lang", Author: "P M F", ESBN: "ESBN-2353-1422"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Go Lang","Author":"P M F","ESBN":"ESBN-2353-1422"}`, string(json), "Book Json marshling wrong")

}

func TestBookFromJson(t *testing.T) {
	json := []byte(`{"Title":"Go Lang","Author":"P M F","ESBN":"ESBN-2353-1422"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Go Lang", Author: "P M F", ESBN: "ESBN-2353-1422"}, book, "Book Json unmarshling wrong")

}
