package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Data struct {
	Size   int64  `json:"size"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func (box *MainBox) AddItem(item Item) []Item {
	box.Item = append(box.Item, item)
	return box.Item
}

type MainBox struct {
	Item []Item `json:"file"`
}
type Item struct {
	Size     int64  `json:"size"`
	Name     string `json:"name"`
	IsDir    bool   `json:"isdir"`
	FullPath string `json:"full_path"`
}

func ls(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("p")
	nmw := ""
	if name == "" {
		name = "./file"
	} else {
		if string(name[0]) != "/" {
			nmw = name
			name = "./file/" + name
		} else {
			nmw = name
			name = "./file" + name
		}
	}
	go fmt.Println(name)
	files, err := ioutil.ReadDir(name)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
	}
	ItemList := []Item{}
	box := MainBox{ItemList}
	run := []rune(nmw)
	slash := "/"
	if string(run[len(run)-1:]) == "/" {
		slash = ""

	}
	for _, f := range files {

		go fmt.Println(f.Name())
		//runes := []rune(f.Name())
		box.AddItem(Item{
			Size:     f.Size(),
			Name:     f.Name(),
			IsDir:    f.IsDir(),
			FullPath: nmw + slash + f.Name(),
		})
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(box)
}
func Status(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("p")
	if name == "" {
		name = "./file"
	} else {
		if string(name[0]) != "/" {
			name = "./file/" + name
		} else {
			name = "./file" + name
		}
	}
	fmt.Println(name)
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
	}
	defer file.Close()

	// получить размер файла
	check := true
	stat, err := file.Stat()
	var size int64
	if os.IsNotExist(err) {
		check = false
	}
	if check == true {
		size = stat.Size()
	} else {
		size = 0
		fmt.Println(size)
	}
	runes := []rune(name)

	d := &Data{
		Size:   size,
		Name:   string(runes[7:]),
		Status: check,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
	//	fmt.Fprint(w, name, stat)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/stat", Status)
	router.HandleFunc("/ls", ls)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":9999", nil)
}

//dd
