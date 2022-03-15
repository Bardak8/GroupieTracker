package main

type Film struct {
	Res []ViewData
}

type People struct {
	Res1 []Viewpeople
}

type Location struct {
	Res2 []ViewLocation
}

type Species struct {
	Res3 []ViewSpecies
}

type Vehicles struct {
	Res4 []ViewVehicles
}

type ViewSpecies struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Classification string   `json:"classification"`
	EyeColors      string   `json:"eye_colors"`
	HairColors     string   `json:"hair_colors"`
	People         []string `json:"people"`
	Films          []string `json:"films"`
	URL            string   `json:"url"`
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

type ViewVehicles struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	VehicleClass string   `json:"vehicle_class"`
	Length       string   `json:"length"`
	Pilot        string   `json:"pilot"`
	Films        []string `json:"films"`
	URL          string   `json:"url"`
}
