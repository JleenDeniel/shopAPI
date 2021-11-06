package main 

import (
	"net/http"
	"log"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Запуск веб-сервера на http://127.0.0.1:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет!"))
}