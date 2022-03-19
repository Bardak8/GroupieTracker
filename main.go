package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//func items(id) {
//	for i := 0; i < len(ViewDataLocation.Res2); i++ {
//		Locationtemp.Films[i] == "descrmovie/"+id
//	}
//}

func main() {
	viewData := Film{Res: loadAPI()}
	viewDataPers := People{Res1: loadAPIpeople()}
	ViewDataLocation := Location{Res2: loadAPILocation1()}
	ViewDataSpecies := Species{Res3: loadAPISpecies()}
	ViewDataVehicles := Vehicles{Res4: loadAPIVehicles()}
	//viewData := loadAPI()
	tmplpage := template.Must(template.ParseFiles("page/page1.html"))
	tmplmovies := template.Must(template.ParseFiles("page/movies.html"))
	tmpldscrmovies := template.Must(template.ParseFiles("page/descrmovie.html"))
	tmplpeople := template.Must(template.ParseFiles("page/people.html"))
	tmpldscrpeople := template.Must(template.ParseFiles("page/descrpeople.html"))
	tmpllocation := template.Must(template.ParseFiles("page/location.html"))
	tmpldscrlocation := template.Must(template.ParseFiles("page/descrlocation.html"))
	tmplspecies := template.Must(template.ParseFiles("page/species.html"))
	tmpldscrspecies := template.Must(template.ParseFiles("page/descrspecies.html"))
	tmplvehicles := template.Must(template.ParseFiles("page/vehicles.html"))
	tmpldscrvehicles := template.Must(template.ParseFiles("page/descrvehicles.html"))

	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplpage.Execute(w, viewData)
	})

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		tmplpeople.Execute(w, viewDataPers)
	})

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		tmplmovies.Execute(w, viewData)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		tmpllocation.Execute(w, ViewDataLocation)
	})

	http.HandleFunc("/species", func(w http.ResponseWriter, r *http.Request) {
		tmplspecies.Execute(w, ViewDataSpecies)
	})

	http.HandleFunc("/vehicles", func(w http.ResponseWriter, r *http.Request) {
		tmplvehicles.Execute(w, ViewDataVehicles)
	})

	http.HandleFunc("/descrmovie/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrmovie/", "")
		viewData1 := loadAPI2(id)
		tmpldscrmovies.Execute(w, viewData1)
	})

	http.HandleFunc("/descrpeople/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrpeople/", "")
		people := loadAPIpeople1(id)
		films := Film{Res: []ViewData{}}

		for _, url := range people.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}

		type Resultat struct {
			Peoples Viewpeople
			Films   Film
		}

		res := Resultat{
			people,
			films,
		}
		tmpldscrpeople.Execute(w, res)
	})

	http.HandleFunc("/descrlocation/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrlocation/", "")
		location := loadAPILocation2(id)
		films := Film{Res: []ViewData{}}

		for _, url := range location.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}

		type Resultat struct {
			Location ViewLocation
			Films    Film
		}

		res := Resultat{
			location,
			films,
		}

		tmpldscrlocation.Execute(w, res)
	})

	http.HandleFunc("/descrspecies/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrspecies/", "")
		viewDataSpecies1 := loadAPISpecies1(id)
		tmpldscrspecies.Execute(w, viewDataSpecies1)
	})

	http.HandleFunc("/descrvehicles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrvehicles/", "")
		vehicles := loadAPIVehicles1(id)
		films := Film{Res: []ViewData{}}

		for _, url := range vehicles.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}

		type Resultat struct {
			Vehicles ViewVehicles
			Films    Film
		}

		res := Resultat{
			vehicles,
			films,
		}

		tmpldscrvehicles.Execute(w, res)
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
