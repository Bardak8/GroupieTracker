package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ViewData []struct {
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

func loadAPI() ViewData {
	var vd ViewData

	url := "https://ghibliapi.herokuapp.com/films"

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

func main() {
	viewData := loadAPI()
	tmplpage1 := template.Must(template.ParseFiles("page/page1.html"))
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		tmplpage1.Execute(w, viewData)
	})
	fmt.Println("Starting server on port 80")
	http.ListenAndServe(":80", nil)
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
