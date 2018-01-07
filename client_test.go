package owm

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWeatherByCityName(t *testing.T) {
	expect := OpenWeatherMap{
		Coord: Coord{Lon: -97.74, Lat: 30.27},
		Weather: []Weather{
			{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01d",
			},
		},
		Main: Main{
			Temp:     296.82,
			Pressure: 1012,
			Humidity: 25,
			TempMin:  296.15,
			TempMax:  298.15,
		},
		Visibility: 11265,
		Wind: Wind{
			Speed: 6.2,
			Deg:   200,
			Gust:  9.3,
		}, Clouds: Clouds{All: 1},
		DT: Time(time.Unix(1511561700, 0)),
		Sys: Sys{
			Country: "US",
			Sunrise: Time(time.Unix(1511528705, 0)),
			Sunset:  Time(time.Unix(1511566248, 0)),
		},
		ID:   4671654,
		Name: "Austin",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"coord":{"lon":-97.74,"lat":30.27},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":296.82,"pressure":1012,"humidity":25,"temp_min":296.15,"temp_max":298.15},"visibility":11265,"wind":{"speed":6.2,"deg":200,"gust":9.3},"clouds":{"all":1},"dt":1511561700,"sys":{"type":1,"id":2557,"message":0.1722,"country":"US","sunrise":1511528705,"sunset":1511566248},"id":4671654,"name":"Austin","cod":200}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := New("Test")
	client.BaseURL = server.URL

	got, err := client.ByCityName("Austin")
	if err != nil {
		t.Error(err)
	}

	t.Logf("got %#v expect %#v\n", got, expect)
}

func TestWeatherByZip(t *testing.T) {
	expect := OpenWeatherMap{
		Coord: Coord{Lon: -97.74, Lat: 30.27},
		Weather: []Weather{
			{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01d",
			},
		},
		Main: Main{
			Temp:     296.82,
			Pressure: 1012,
			Humidity: 25,
			TempMin:  296.15,
			TempMax:  298.15,
		},
		Visibility: 11265,
		Wind: Wind{
			Speed: 6.2,
			Deg:   200,
			Gust:  9.3,
		}, Clouds: Clouds{All: 1},
		DT: Time(time.Unix(1511561700, 0)),
		Sys: Sys{
			Country: "US",
			Sunrise: Time(time.Unix(1511528705, 0)),
			Sunset:  Time(time.Unix(1511566248, 0)),
		},
		ID:   4671654,
		Name: "Austin",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"coord":{"lon":-97.74,"lat":30.27},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":296.82,"pressure":1012,"humidity":25,"temp_min":296.15,"temp_max":298.15},"visibility":11265,"wind":{"speed":6.2,"deg":200,"gust":9.3},"clouds":{"all":1},"dt":1511561700,"sys":{"type":1,"id":2557,"message":0.1722,"country":"US","sunrise":1511528705,"sunset":1511566248},"id":4671654,"name":"Austin","cod":200}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := New("Test")
	client.BaseURL = server.URL

	got, err := client.ByZip("78704")
	if err != nil {
		t.Error(err)
	}

	t.Logf("got %#v expect %#v\n", got, expect)
}

func TestWeatherByCoordinates(t *testing.T) {
	expect := OpenWeatherMap{
		Coord: Coord{Lon: -97.74, Lat: 30.27},
		Weather: []Weather{
			{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01d",
			},
		},
		Main: Main{
			Temp:     296.82,
			Pressure: 1012,
			Humidity: 25,
			TempMin:  296.15,
			TempMax:  298.15,
		},
		Visibility: 11265,
		Wind: Wind{
			Speed: 6.2,
			Deg:   200,
			Gust:  9.3,
		}, Clouds: Clouds{All: 1},
		DT: Time(time.Unix(1511561700, 0)),
		Sys: Sys{
			Country: "US",
			Sunrise: Time(time.Unix(1511528705, 0)),
			Sunset:  Time(time.Unix(1511566248, 0)),
		},
		ID:   4671654,
		Name: "Austin",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"coord":{"lon":-97.74,"lat":30.27},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":296.82,"pressure":1012,"humidity":25,"temp_min":296.15,"temp_max":298.15},"visibility":11265,"wind":{"speed":6.2,"deg":200,"gust":9.3},"clouds":{"all":1},"dt":1511561700,"sys":{"type":1,"id":2557,"message":0.1722,"country":"US","sunrise":1511528705,"sunset":1511566248},"id":4671654,"name":"Austin","cod":200}`)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := New("Test")
	client.BaseURL = server.URL

	got, err := client.ByCoordinates(30.27, -97.74)
	if err != nil {
		t.Error(err)
	}

	t.Logf("got %#v expect %#v\n", got, expect)
}
