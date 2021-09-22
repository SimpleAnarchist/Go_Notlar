package dersler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func GetirGet(){
	response,err := http.Get("https://jsonplaceholder.typicode.com/todos/1") //hata var mı diye 2. değişken err
	//https://pkg.go.dev/net/http
	if err != nil{
		fmt.Println(err)
	}
	defer response.Body.Close() //bitince kapat
	//http tipinde geldiği için dönüşümler
	bodyBytes,_ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes,&todo)  //json dan go dönüştürme
	fmt.Println(todo)
}

func GonderPost(){
	todo := Todo{1,2,"Alışveriş Yapılacak", false}
	jsonTodo,err := json.Marshal(todo) //go da json dönüştürme
	if err != nil{
		fmt.Println(err)
	}
	response, err := http.Post(
		"https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))
	if err != nil{
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes,_ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes,&todoResponse)  //json dan go dönüştürme
	fmt.Println(todo)
}