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
		films := loadAPI2(id)
		peoples := People{Res1: []Viewpeople{}}
		species := Species{Res3: []ViewSpecies{}}
		locations := Location{Res2: []ViewLocation{}}
		vehicles := Vehicles{Res4: []ViewVehicles{}}

		for _, url := range films.People {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/people/", "")
			if id == "" {
				peoples.Res1 = []Viewpeople{}
			} else {
				people := loadAPIpeople1(id)
				peoples.Res1 = append(peoples.Res1, people)
			}
		}

		for _, url := range films.Species {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/species/", "")
			if id == "" {
				species.Res3 = []ViewSpecies{}
			} else {
				specie := loadAPISpecies1(id)
				species.Res3 = append(species.Res3, specie)
			}
		}

		for _, url := range films.Vehicles {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/vehicles/", "")
			if id == "" {
				vehicles.Res4 = []ViewVehicles{}
			} else {
				vehicle := loadAPIVehicles1(id)
				vehicles.Res4 = append(vehicles.Res4, vehicle)
			}
		}

		type Resultat struct {
			Films     ViewData
			People    People
			Species   Species
			Locations Location
			Vehicles  Vehicles
		}
		res := Resultat{
			films,
			peoples,
			species,
			locations,
			vehicles,
		}
		tmpldscrmovies.Execute(w, res)
	})

	http.HandleFunc("/descrpeople/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrpeople/", "")
		people := loadAPIpeople1(id)
		films := Film{Res: []ViewData{}}
		species := loadAPISpeciesFullUrl(people.Species)
		for _, url := range people.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}

		type Resultat struct {
			Peoples Viewpeople
			Films   Film
			Species ViewSpecies
		}

		res := Resultat{
			people,
			films,
			species,
		}
		tmpldscrpeople.Execute(w, res)
	})

	http.HandleFunc("/descrlocation/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrlocation/", "")
		location := loadAPILocation2(id)
		films := Film{Res: []ViewData{}}
		peoples := People{Res1: []Viewpeople{}}

		for _, url := range location.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}

		for _, url := range location.Residents {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			people := loadAPIpeople1(id)
			peoples.Res1 = append(peoples.Res1, people)
		}

		type Resultat struct {
			Location ViewLocation
			Films    Film
			People   People
		}

		res := Resultat{
			location,
			films,
			peoples,
		}

		tmpldscrlocation.Execute(w, res)
	})

	http.HandleFunc("/descrspecies/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrspecies/", "")
		species := loadAPISpecies1(id)
		films := Film{Res: []ViewData{}}
		peoples := People{Res1: []Viewpeople{}}

		for _, url := range species.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}

		for _, url := range species.People {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/people/", "")
			people := loadAPIpeople1(id)
			peoples.Res1 = append(peoples.Res1, people)
		}

		type Resultat struct {
			Species ViewSpecies
			Films   Film
			People  People
		}

		res := Resultat{
			species,
			films,
			peoples,
		}
		tmpldscrspecies.Execute(w, res)
	})

	http.HandleFunc("/descrvehicles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.ReplaceAll(r.URL.Path, "/descrvehicles/", "")
		vehicles := loadAPIVehicles1(id)
		films := Film{Res: []ViewData{}}
		peoples := loadAPIPeopleFullUrl(vehicles.Pilot)
		for _, url := range vehicles.Films {
			id := strings.ReplaceAll(url, "https://ghibliapi.herokuapp.com/films/", "")
			film := loadAPI2(id)
			films.Res = append(films.Res, film)
		}
		type Resultat struct {
			Vehicles ViewVehicles
			Films    Film
			People   Viewpeople
		}
		res := Resultat{
			vehicles,
			films,
			peoples,
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
