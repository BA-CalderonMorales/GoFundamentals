package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())
	return lines
}

func EnglishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Internet!")
}

func SpanishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hola Internet!")
}

func InteractHandler(writer http.ResponseWriter, request *http.Request) {
	todoVals := getStrings("todos.txt")
	fmt.Printf("%#v\n", todoVals)
	tmpl, err := template.ParseFiles("index.html")
	errorCheck(err)
	todos := ToDoList{
		ToDoCount: len(todoVals),
		ToDos:     todoVals,
	}
	err = tmpl.Execute(writer, todos)
}

func NewHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("new.html")
	errorCheck(err)
	err = tmpl.Execute(writer, nil)
}

func CreateHandler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	err = file.Close() // can use interchangeably with defer
	errorCheck(err)
	http.Redirect(writer, request, "/interact", http.StatusFound)
}

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/hello", EnglishHandler)
	http.HandleFunc("/hola", SpanishHandler)
	http.HandleFunc("/interact", InteractHandler)
	http.HandleFunc("/new", NewHandler)
	http.HandleFunc("/create", CreateHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
