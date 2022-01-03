package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// var athelet = "https://allsports.tw/view/777889/777889/zekken/11886/"
// var athelet = "http://allsports.tw/ajax/photo/list/777922/777922/zekken/20203/0/"
// var athelet = "https://allsports.tw/ajax/photo/list/777994/777994/zekken/3165/0/"
// var athelet = "https://allsports.tw/ajax/photo/list/777994/777994/zekken/2668/0/"
var athelet = "https://allsports.tw/ajax/photo/list/778094/778094/zekken/11436/0/"

// var athelet = "http://allsports.tw/ajax/photo/list/777889/777889/zekken/11887/0/"
var savepath = "c://user//temp//"

type ath struct {
	Id        string `json:"id"`
	forsale   bool
	ss        bool
	time      string
	params    paramtype
	Image     string `json:"Image"`
	Detail    string `json:"Detail"`
	iscampain bool
}

type paramtype struct {
	event_id         string
	photographer_id  string
	photo_num        string
	photo_id         string
	page_id          string
	photo_hash       string
	album_id         string
	image_cache_flag bool
	logo_select      bool
	imageuritype     string
	logo_id          int
	watermark_id     int
}

func httpallsp() {
	resp, err := http.Get(athelet)

	if err != nil {
		log.Println("http get error:", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("body read error:", err)
	}

	// bodystr := string(body)

	// fmt.Println(bodystr)

	var a []ath
	// err = json.Unmarshal([]byte(bodystr), &a)
	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("a = ", a, len(a))
	fmt.Println("a[0] = ", a[0])
	fmt.Println("a[0].id = ", a[0].Id, a[0].Image, a[0].Detail)
	var filename string
	for i := 0; i < len(a); i++ {
		fmt.Println("wget " + a[i].Detail)
		filename = savepath + strconv.Itoa(i) + ".jpg"
		err = DownloadFile(filename, a[i].Detail)
		if err != nil {
			panic(err)
		}
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
