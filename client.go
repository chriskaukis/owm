package owm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	version          = "0.1"
	defaultBaseURL   = "https://api.openweathermap.org/data/2.5"
	defaultUserAgent = "bike-barometer/" + version
	defaultTimeout   = time.Second * 60
)

type Coord struct {
	Lon float64 `json:"lon,omitempty"`
	Lat float64 `json:"lat,omitempty"`
}

type Weather struct {
	ID          int    `json:"id,omitempty"`
	Main        string `json:"main,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
}

type Main struct {
	Temp     float64 `json:"temp,omitempty"`
	Pressure int     `json:"pressure,omitempty"`
	Humidity int     `json:"humidity,omitempty"`
	TempMin  float64 `json:"temp_min,omitempty"`
	TempMax  float64 `json:"temp_max,omitempty"`
}

type Wind struct {
	Speed float64 `json:"speed,omitempty"`
	Deg   int     `json:"deg,omitempty"`
	Gust  float64 `json:"gust,omitempty"`
}

type Clouds struct {
	All int `json:"all,omitempty"`
}

type Rain struct {
	VolumeForLast3Hours int `json:"3h,omitempty"`
}

type Snow struct {
	VolumeForLast3Hours int `json:"3h,omitempty"`
}

type Sys struct {
	Country string `json:"country,omitempty"`
	Sunrise Time   `json:"sunrise,omitempty"`
	Sunset  Time   `json:"sunset,omitempty"`
}
type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	timestamp, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Time(time.Unix(int64(timestamp), 0))
	return nil
}

type OpenWeatherMap struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	DT         Time   `json:"dt,omitempty"`
	Visibility int    `json:"visibility,omitempty"`
	Coord      `json:"coord,omitempty"`
	Weather    []Weather `json:"weather,omitempty"`
	Main       `json:"main,omitempty"`
	Wind       `json:"wind,omitempty"`
	Snow       `json:"snow,omitempty"`
	Clouds     `json:"clouds,omitempty"`
	Sys        `json:"sys,omitempty"`
}

// OpenWeatherMapper represents an interface of actions that we currently support using the OpenWeatherMap API.
type OpenWeatherMapper interface {
	ByCityName(string) (*OpenWeatherMap, error)
	ByCoordinates(int, int) (*OpenWeatherMap, error)
	ByZipCode(string) (*OpenWeatherMap, error)
}

// Client represents an OpenWeatherMap API client. Client implements the OpenWeatherMappper interface.
type Client struct {
	*http.Client
	BaseURL   string
	UserAgent string
	key       string
}

// New returns a new client with the absolute minimum information needed to interact with the OpenWeatherMap API.
func New(key string) *Client {
	return &Client{
		Client:    &http.Client{Timeout: defaultTimeout},
		BaseURL:   defaultBaseURL,
		UserAgent: defaultUserAgent,
		key:       key,
	}
}

func encodeJSON(body interface{}) (io.ReadWriter, error) {
	if body == nil {
		return nil, nil
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	return buf, err
}

// Request makes a request to the OpenWeatherMap API.
func (c *Client) Request(method, path string, body interface{}) (*http.Request, error) {
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	rel := &url.URL{Path: path}
	u := baseURL.ResolveReference(rel)

	buf, err := encodeJSON(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	return res, err
}

// ByCityName gets and returns the weather data or an error for the given city name.
// http://openweathermap.org/current#name
// curl "http://api.openweathermap.org/data/2.5/weather?q=Austin&APPID=8c0ada06102f197d47234e491265501b"
func (c *Client) ByCityName(city string) (*OpenWeatherMap, error) {
	req, err := c.Request(http.MethodGet, "/weather", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("q", city)
	query.Add("APPID", c.key)
	req.URL.RawQuery = query.Encode()

	var weather OpenWeatherMap
	_, err = c.Do(req, &weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}

func (c *Client) ByZip(zip string) (*OpenWeatherMap, error) {
	req, err := c.Request(http.MethodGet, "/weather", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("zip", zip)
	query.Add("APPID", c.key)
	req.URL.RawQuery = query.Encode()

	var weather OpenWeatherMap
	_, err = c.Do(req, &weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}

func (c *Client) ByCoordinates(lat, lon float64) (*OpenWeatherMap, error) {
	req, err := c.Request(http.MethodGet, "/weather", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("lat", fmt.Sprintf("%f", lat))
	query.Add("lon", fmt.Sprintf("%f", lon))
	query.Add("APPID", c.key)
	req.URL.RawQuery = query.Encode()

	var weather OpenWeatherMap
	_, err = c.Do(req, &weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}
