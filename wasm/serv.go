package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
	//fmt.Fprint(w, "index")
}
func JsRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	object := vars["object"]
	//response := fmt.Sprintf("%s", id)
	path := "./js/" + object
	fmt.Println(path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, path)
}
func wasmRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	object := vars["object"]
	//response := fmt.Sprintf("%s", id)
	path := "./wasm/" + object
	fmt.Println(path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	http.ServeFile(w, r, path)
}
func main() {
	router := mux.NewRouter()
	//file routers
	router.HandleFunc("/js/{object}", JsRouter)
	router.HandleFunc("/wasm/{object}", wasmRouter)
	router.HandleFunc("/", index)

	//run serve
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
