package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

type Bloup struct {
	Res []ViewData
}

type Bloup1 struct {
	Res1 []Viewpeople
}

type Bloup2 struct {
	Res2 []ViewLocation
}

type Viewpeople struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender,omitempty"`
	Age       string   `json:"age"`
	EyeColor  string   `json:"eye_color"`
	HairColor string   `json:"hair_color"`
	Films     []string `json:"films"`
	Species   string   `json:"species"`
	URL       string   `json:"url"`
	Gander    string   `json:"gander,omitempty"`
}

type ViewData struct {
	ID                     string   `json:"id"`
	Title                  string   `json:"title"`
	OriginalTitle          string   `json:"original_title"`
	OriginalTitleRomanised string   `json:"original_title_romanised"`
	Image                  string   `json:"image"`
	MovieBanner            string   `json:"movie_banner"`
	Description            string   `json:"description"`
	Director               string   `json:"director"`
	Producer               string   `json:"producer"`
	ReleaseDate            string   `json:"release_date"`
	RunningTime            string   `json:"running_time"`
	RtScore                string   `json:"rt_score"`
	People                 []string `json:"people"`
	Species                []string `json:"species"`
	Locations              []string `json:"locations"`
	Vehicles               []string `json:"vehicles"`
	URL                    string   `json:"url"`
}

type ViewLocation struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Climate      string   `json:"climate"`
	Terrain      string   `json:"terrain"`
	SurfaceWater string   `json:"surface_water"`
	Residents    []string `json:"residents"`
	Films        []string `json:"films"`
	URL          string   `json:"url"`
}

func loadAPI() []ViewData {
	var vd []ViewData

	url := "https://ghibliapi.herokuapp.com/films/"

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vd
}

func loadAPI2(id string) ViewData {
	var vd ViewData

	url := "https://ghibliapi.herokuapp.com/films/" + id

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		debug.PrintStack()
		log.Fatal(jsonErr)
	}

	return vd
}

func loadAPIpeople() []Viewpeople {
	var vd []Viewpeople

	url := "https://ghibliapi.herokuapp.com/people/"

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vd
}

func loadAPIpeople1(id string) Viewpeople {
	var vd Viewpeople

	url := "https://ghibliapi.herokuapp.com/people/" + id

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		debug.PrintStack()
		log.Fatal(jsonErr)
	}

	return vd
}

func loadAPILocation1() []ViewLocation {
	var vd []ViewLocation

	url := "https://ghibliapi.herokuapp.com/locations/"

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vd
}

func loadAPILocation2(id string) ViewLocation {
	var vd ViewLocation

	url := "https://ghibliapi.herokuapp.com/locations/" + id

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		debug.PrintStack()
		log.Fatal(jsonErr)
	}

	return vd
}

func main() {
	viewData := Bloup{Res: loadAPI()}
	viewDatapers := Bloup1{Res1: loadAPIpeople()}
	ViewDatalocation := Bloup2{Res2: loadAPILocation1()}
	//viewData := loadAPI()
	tmplpage := template.Must(template.ParseFiles("page/page1.html"))
	tmplpage1 := template.Must(template.ParseFiles("page/movies.html"))
	tmplpage2 := template.Must(template.ParseFiles("page/descrmovie.html"))
	tmplpage4 := template.Must(template.ParseFiles("page/people.html"))
	tmplpage3 := template.Must(template.ParseFiles("page/descrpeople.html"))
	tmplpage5 := template.Must(template.ParseFiles("page/location.html"))
	tmplpage6 := template.Must(template.ParseFiles("page/descrlocation.html"))

	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplpage.Execute(w, viewData)
	})

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		tmplpage4.Execute(w, viewDatapers)
	})

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		tmplpage1.Execute(w, viewData)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		tmplpage5.Execute(w, ViewDatalocation)
	})

	http.HandleFunc("/descrmovie/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrmovie/", "")
		viewData1 := loadAPI2(id)
		tmplpage2.Execute(w, viewData1)
	})

	http.HandleFunc("/descrpeople/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrpeople/", "")
		viewDatapers1 := loadAPIpeople1(id)
		tmplpage3.Execute(w, viewDatapers1)
	})

	http.HandleFunc("/descrlocation/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrlocation/", "")
		viewDataLocation1 := loadAPILocation2(id)
		tmplpage6.Execute(w, viewDataLocation1)
	})

	fmt.Println("Starting server on port 80")
	http.ListenAndServe(":8000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

/*
func handleghibli() {
	var responses []Response
	//TO DO Implement handler
	response, err := http.Get("https://ghibliapi.herokuapp.com/films/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responses)
	for i := range responses {
		fmt.Println(responses[i].Id)
	}
} */
