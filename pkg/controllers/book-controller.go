package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../../models"
	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetBookbyId(w http.ResponseWriter, r *http.Request) {

	Id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 0, 0)
	if err != nil {
		fmt.Println("parse error")
	}
	bookdetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newbook := &models.Book{}
	ParseBody(r, newbook)
	b := newbook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 0, 0)
	if err != nil {
		fmt.Println("parse error")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(mux.Vars(r)["bookId"], 0, 0)
	if err != nil {
		fmt.Println("parse error")
	}
	Updated := &models.Book{}
	ParseBody(r, Updated)
	bookDetails, db := models.GetBookById(Id)
	if Updated.Name != "" {
		bookDetails.Name = Updated.Name
	}
	if Updated.Author != "" {
		bookDetails.Author = Updated.Author
	}
	if Updated.Publication != "" {
		bookDetails.Publication = Updated.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
