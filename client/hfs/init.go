package hfs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (api *APIStruct) Hello() string {
	return "hello"
}

type MainBox struct {
	Item []Item `json:"file"`
}
type Item struct {
	Size  int64  `json:"size"`
	Name  string `json:"name"`
	IsDir bool   `json:"isdir"`
}

func (api *APIStruct) Ls(names ...string) MainBox {
	file := ""
	x := 0
	url := "http://"
	for _, n := range names {
		file = file + "f=" + n
		//fmt.Printf("Hello %s\n", n)
		x++
	}
	//	fmt.Println(file)
	//	fmt.Println(x)
	if x == 0 {
		url = url + api.Url + "/ls"
	} else {
		url = url + api.Url + "/ls?p=" + file
	}
	//	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	//	bod := string(body)
	//	fmt.Println(bod)
	var s MainBox
	err = json.Unmarshal(body, &s)
	if err != nil {
		print("whoops:", err)
	}
	return s
}
func (api *APIStruct) Stat(name string) string {
	resp, err := http.Get("http://" + api.Url + "/stat?p=" + name)
	if err != nil {
		print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	reqv := string(body)
	return reqv
}

type APIStruct struct {
	Url string
}

func API(url string) *APIStruct {
	return &APIStruct{
		Url: url,
	}
}
